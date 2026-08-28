import 'package:flutter_test/flutter_test.dart';
import 'package:op_memory_game/models/card_model.dart';
import 'package:op_memory_game/providers/game_provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() {
  // Keeps every provider created by a test so tearDown can stop its timer.
  final createdGames = <GameProvider>[];

  setUp(() {
    // Fresh in-memory preferences for every test (best score isolation).
    SharedPreferences.setMockInitialValues({});
  });

  tearDown(() {
    for (final game in createdGames) {
      game.dispose(); // cancels the game timer
    }
    createdGames.clear();
  });

  /// Creates a provider with a short flip-back delay (so the mismatch test
  /// does not wait in real time) and lets the constructor's async
  /// preferences load complete before the test body runs.
  Future<GameProvider> createGame({
    Duration flipBackDelay = const Duration(milliseconds: 10),
  }) async {
    final game = GameProvider(flipBackDelay: flipBackDelay);
    createdGames.add(game);
    await Future<void>.delayed(Duration.zero);
    return game;
  }

  /// Returns the index of the only other card showing the same image as
  /// the card at [index] (each image is dealt exactly twice).
  int partnerOf(List<CardModel> cards, int index) {
    for (var i = 0; i < cards.length; i++) {
      if (i != index && cards[i].imagePath == cards[index].imagePath) {
        return i;
      }
    }
    fail('No partner card found for index $index');
  }

  test('initial state deals 16 face-down cards in 8 pairs, game not over',
      () async {
    final game = await createGame();

    expect(game.cards.length, 16);
    expect(game.isGameOver, isFalse);
    expect(game.moves, 0);
    expect(game.secondsElapsed, 0);
    expect(game.flippedIndices, isEmpty);
    for (final card in game.cards) {
      expect(card.isFlipped, isFalse,
          reason: '${card.imagePath} should start face-down');
      expect(card.isMatched, isFalse);
    }
    final counts = <String, int>{};
    for (final card in game.cards) {
      counts[card.imagePath] = (counts[card.imagePath] ?? 0) + 1;
    }
    expect(counts.length, 8);
    expect(counts.values.every((count) => count == 2), isTrue);
  });

  test('flipping a card turns it face-up', () async {
    final game = await createGame();

    game.onCardTap(0);

    expect(game.cards[0].isFlipped, isTrue);
    expect(game.flippedIndices, [0]);
  });

  test('matching two identical cards marks both as matched', () async {
    final game = await createGame();

    final second = partnerOf(game.cards, 0);
    game.onCardTap(0);
    game.onCardTap(second);

    expect(game.cards[0].isMatched, isTrue);
    expect(game.cards[second].isMatched, isTrue);
    // Matched cards stay face-up and leave the comparison queue.
    expect(game.cards[0].isFlipped, isTrue);
    expect(game.flippedIndices, isEmpty);
  });

  test('non-matching pair flips back after the delay and blocks extra taps',
      () async {
    final game = await createGame(flipBackDelay: const Duration(milliseconds: 10));

    final other =
        game.cards.indexWhere((card) => card.imagePath != game.cards[0].imagePath);
    game.onCardTap(0);
    game.onCardTap(other);
    expect(game.moves, 1);

    // While two cards are face-up, tapping a third one is ignored.
    final third = [
      for (var i = 0; i < game.cards.length; i++)
        if (i != 0 && i != other) i,
    ].first;
    game.onCardTap(third);
    expect(game.cards[third].isFlipped, isFalse);
    expect(game.flippedIndices.length, 2);

    // After the flip-back delay both cards are face-down again.
    await Future<void>.delayed(const Duration(milliseconds: 100));
    expect(game.cards[0].isFlipped, isFalse);
    expect(game.cards[other].isFlipped, isFalse);
    expect(game.flippedIndices, isEmpty);
  });

  test('move counter increments once per pair attempt', () async {
    final game = await createGame();

    game.onCardTap(0);
    game.onCardTap(partnerOf(game.cards, 0));
    expect(game.moves, 1);

    final next = game.cards.indexWhere((card) => !card.isMatched);
    game.onCardTap(next);
    game.onCardTap(partnerOf(game.cards, next));
    expect(game.moves, 2);
  });

  test('matching every pair sets the game-over flag and best score', () async {
    final game = await createGame();

    while (game.cards.any((card) => !card.isMatched)) {
      final first = game.cards.indexWhere((card) => !card.isMatched);
      game.onCardTap(first);
      game.onCardTap(partnerOf(game.cards, first));
    }

    expect(game.isGameOver, isTrue);
    expect(game.moves, 8);
    expect(game.calculateScore(), greaterThan(0));
    expect(game.bestScore, game.calculateScore());
    // Let the async preference write finish before the test ends.
    await Future<void>.delayed(Duration.zero);
  });

  test('restart returns the game to its initial state', () async {
    final game = await createGame();

    game.onCardTap(0);
    game.onCardTap(partnerOf(game.cards, 0));
    expect(game.moves, greaterThan(0));

    game.initializeGame();

    expect(game.moves, 0);
    expect(game.secondsElapsed, 0);
    expect(game.isGameOver, isFalse);
    expect(game.flippedIndices, isEmpty);
    expect(game.cards.length, 16);
    expect(
      game.cards.every((card) => !card.isFlipped && !card.isMatched),
      isTrue,
    );
  });
}
