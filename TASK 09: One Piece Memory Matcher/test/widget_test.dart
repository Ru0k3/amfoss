import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:op_memory_game/main.dart';
import 'package:op_memory_game/providers/game_provider.dart';
import 'package:op_memory_game/widgets/memory_card.dart';
import 'package:provider/provider.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() {
  TestWidgetsFlutterBinding.ensureInitialized();

  setUp(() {
    // Fresh in-memory preferences for every test (best score isolation).
    SharedPreferences.setMockInitialValues({});
  });

  /// Pumps [MyApp] under a provider, exactly like [main] does. The provider
  /// is created by the test (not by the widget tree) so it can be disposed
  /// deterministically and no game timer outlives the test.
  Future<GameProvider> pumpApp(WidgetTester tester) async {
    final game = GameProvider();
    await tester.pumpWidget(
      ChangeNotifierProvider<GameProvider>.value(
        value: game,
        child: const MyApp(),
      ),
    );
    await tester.pump(); // let the async best-score load settle
    return game;
  }

  /// Reads which image the given card currently shows (front or back).
  String shownAssetOf(WidgetTester tester, Finder card) {
    final container = tester.widget<Container>(
      find.descendant(of: card, matching: find.byType(Container)).first,
    );
    final decoration = container.decoration! as BoxDecoration;
    return (decoration.image!.image as AssetImage).assetName;
  }

  testWidgets('app builds and renders the 4x4 card grid', (tester) async {
    final game = await pumpApp(tester);

    expect(find.text('One Piece Memory Matcher'), findsOneWidget);
    expect(find.byType(GridView), findsOneWidget);
    expect(find.byType(MemoryCard), findsNWidgets(16));
    expect(find.text('Best Score'), findsOneWidget);
    expect(find.text('Time'), findsOneWidget);
    expect(find.text('Moves'), findsOneWidget);
    expect(find.text('Restart Game'), findsOneWidget);

    game.dispose();
  });

  testWidgets('tapping a face-down card flips it face-up', (tester) async {
    final game = await pumpApp(tester);

    final card = find.byType(MemoryCard).first;

    // All cards start face-down, showing the card back.
    expect(shownAssetOf(tester, card), 'assets/images/card_back.png');

    await tester.tap(card);
    await tester.pump(); // process the tap and retarget the flip animation
    await tester.pump(const Duration(milliseconds: 600)); // flip lasts 500ms

    expect(shownAssetOf(tester, card), game.cards[0].imagePath);

    game.dispose();
    await tester.pump();
  });
}
