import 'dart:async';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../constants.dart';
import '../models/card_model.dart';

/// Holds the game state for the memory matching game: the deck of cards,
/// flip/match bookkeeping, the move counter and elapsed time, and the
/// persisted best score.
class GameProvider with ChangeNotifier {
  List<CardModel> cards = [];
  List<int> flippedIndices = [];
  int moves = 0;
  int secondsElapsed = 0;
  int bestScore = 0;
  bool isGameOver = false;
  Timer? _timer;

  /// How long a mismatched pair stays face-up before flipping back.
  /// Tests inject a shorter duration so they don't wait in real time.
  final Duration flipBackDelay;

  final List<String> _imageAssets = [
    'assets/images/luffy.png',
    'assets/images/zoro.png',
    'assets/images/nami.png',
    'assets/images/sanji.png',
    'assets/images/chopper.png',
    'assets/images/robin.png',
    'assets/images/ace.png',
    'assets/images/law.png',
  ];

  GameProvider({this.flipBackDelay = defaultFlipBackDelay}) {
    _loadPreferences();
    initializeGame();
  }

  Future<void> _loadPreferences() async {
    final prefs = await SharedPreferences.getInstance();
    bestScore = prefs.getInt('bestScore') ?? 0;
    notifyListeners();
  }

  void initializeGame() {
    _timer?.cancel();
    cards = [];
    flippedIndices = [];
    moves = 0;
    secondsElapsed = 0;
    isGameOver = false;

    List<String> gameAssets = [..._imageAssets, ..._imageAssets];
    gameAssets.shuffle();

    for (int i = 0; i < gameAssets.length; i++) {
      cards.add(CardModel(id: i, imagePath: gameAssets[i]));
    }
    notifyListeners();
  }

  void startTimer() {
    if (_timer != null && _timer!.isActive) return;
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      secondsElapsed++;
      notifyListeners();
    });
  }

  void onCardTap(int index) {
    if (isGameOver || cards[index].isFlipped || cards[index].isMatched || flippedIndices.length >= 2) {
      return;
    }

    if (moves == 0 && flippedIndices.isEmpty) {
      startTimer();
    }

    cards[index].isFlipped = true;
    flippedIndices.add(index);
    notifyListeners();

    if (flippedIndices.length == 2) {
      moves++;
      _checkMatch();
    }
  }

  void _checkMatch() {
    int firstIndex = flippedIndices[0];
    int secondIndex = flippedIndices[1];

    if (cards[firstIndex].imagePath == cards[secondIndex].imagePath) {
      cards[firstIndex].isMatched = true;
      cards[secondIndex].isMatched = true;
      flippedIndices = [];
      
      if (cards.every((card) => card.isMatched)) {
        _endGame();
      }
    } else {
      Future.delayed(flipBackDelay, () {
        cards[firstIndex].isFlipped = false;
        cards[secondIndex].isFlipped = false;
        flippedIndices = [];
        notifyListeners();
      });
    }
    notifyListeners();
  }

  void _endGame() async {
    isGameOver = true;
    _timer?.cancel();
    
    int currentScore = calculateScore();
    if (currentScore > bestScore) {
      bestScore = currentScore;
      final prefs = await SharedPreferences.getInstance();
      await prefs.setInt('bestScore', bestScore);
    }
    notifyListeners();
  }

  int calculateScore() {
    if (moves == 0) return 0;
    // Score = (pairs * points per pair) - move penalty - time penalty.
    // The pair count is derived from the deck so it never goes stale.
    final int score = (cards.length ~/ 2) * scorePerPair -
        moves * scoreMovePenalty -
        secondsElapsed * scoreTimePenalty;
    return score > 0 ? score : 0;
  }

  String get formattedTime {
    int minutes = secondsElapsed ~/ 60;
    int seconds = secondsElapsed % 60;
    return '${minutes.toString().padLeft(2, '0')}:${seconds.toString().padLeft(2, '0')}';
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }
}
