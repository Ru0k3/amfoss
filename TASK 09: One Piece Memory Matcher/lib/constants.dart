import 'package:flutter/material.dart';

// ── Board layout ─────────────────────────────────────────────────────────----

/// Number of columns in the card grid. The board is always square:
/// 4 columns x 4 rows = 16 cards (8 pairs).
const int gridCrossAxisCount = 4;

/// Spacing in logical pixels between cards, both across and down the grid.
const double gridSpacing = 8.0;

/// Maximum width and height in logical pixels of the game board.
const double maxBoardSize = 520;

/// Maximum width in logical pixels of the stats row above the board.
const double maxStatsRowWidth = 600;

/// Corner radius in logical pixels of each card.
const double cardBorderRadius = 8.0;

/// Perspective factor of the 3D flip effect (smaller = flatter projection).
const double cardPerspective = 0.001;

/// Duration of the card flip animation.
const Duration cardFlipDuration = Duration(milliseconds: 500);

// ── Game rules / scoring ─────────────────────────────────────────────────────

/// Default time a mismatched pair stays face-up before flipping back.
const Duration defaultFlipBackDelay = Duration(milliseconds: 1000);

/// Points awarded per matched pair before penalties are applied.
const int scorePerPair = 100;

/// Points deducted per move (each two-card attempt counts as one move).
const int scoreMovePenalty = 10;

/// Points deducted per elapsed second of play.
const int scoreTimePenalty = 2;

// ── Theming ──────────────────────────────────────────────────────────────────

/// Accent color used for the restart button and the win banner.
const Color accentColor = Colors.orange;

/// Text/icon color on top of the accent-colored restart button.
const Color buttonForegroundColor = Colors.white;

/// Color of the stat labels ("Best Score", "Time", "Moves").
const Color statLabelColor = Colors.grey;

/// Font size in logical pixels of the stat labels.
const double statLabelFontSize = 14;

/// Font size in logical pixels of the stat values.
const double statValueFontSize = 20;
