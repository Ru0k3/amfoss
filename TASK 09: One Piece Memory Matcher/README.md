# One Piece Memory Matcher

A One Piece themed memory card matching game built with Flutter — flip cards on a 4x4 grid and find all 8 character pairs in as few moves as possible.

<!-- TODO: Record a short gameplay GIF (e.g. with `flutter run -d chrome` + a screen recorder, or `asciinema`/`ScreenToGif`) and embed it here as `docs/gameplay.gif`. -->

## Features

- **4x4 memory grid** with 8 One Piece character pairs (Luffy, Zoro, Nami, Sanji, Chopper, Robin, Ace, Law), shuffled on every new game
- **3D card flip animation** using `TweenAnimationBuilder` + `Transform` with perspective; mismatched pairs flip back after a short delay
- **Move counter** — each two-card attempt counts as one move
- **Game timer** — starts on the first flip and displays as `mm:ss`
- **Scoring** — `pairs × 100 − moves × 10 − seconds × 2`, floored at 0
- **Persistent best score** — saved locally with `shared_preferences` when you finish a game, restored on startup
- **Win banner** with the final score once all pairs are matched, plus a restart button
- **Dark Material 3 theme** seeded with the Straw Hat orange

## Architecture

```
lib/
├── main.dart               # App entry point: wraps the app in a ChangeNotifierProvider
├── constants.dart           # Documented layout, gameplay and theming constants
├── models/
│   └── card_model.dart      # CardModel: id, imagePath, isFlipped, isMatched
├── providers/
│   └── game_provider.dart   # GameProvider (ChangeNotifier): deck setup, flip/match
│                            #   logic, moves, timer, scoring, best-score persistence
├── screens/
│   └── game_screen.dart     # GameScreen: stats header, card grid, win banner, restart
└── widgets/
    └── memory_card.dart     # MemoryCard: a single card with the 3D flip animation
```

Each layer has one job: **models** are plain data, the **provider** holds all mutable game state and rules, **screens** compose the layout, and **widgets** render reusable pieces. The state management uses **Provider + `ChangeNotifier`** because it keeps game logic completely out of the widget tree: `GameProvider` notifies listeners on every state change and widgets like `GameScreen` rebuild through `Consumer` — no `setState` plumbing, and the game state survives widget rebuilds because the provider lives above `MaterialApp`.

## Getting Started

Prerequisite: the [Flutter SDK](https://docs.flutter.dev/get-started/install) (3.x, web enabled).

```bash
flutter pub get
flutter run -d chrome
```

## Testing

The test suite covers the game logic (initial deal, flipping, matching, flip-back, move counting, game over, restart) and the UI (grid renders, tapping flips a card):

```bash
flutter test
```
