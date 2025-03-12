import 'package:flutter/material.dart';
import 'colors.dart';

class F1Gradients {
  static const speed = LinearGradient(
    colors: [F1Colors.racingRed, Color(0xFFFF3300), Color(0xFFFF6600)],
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
  );

  static const track = LinearGradient(
    colors: [
      F1Colors.trackBlack,
      Color(0xFF1E1E1E),
      Color(0xFF2A2A2A),
      Color(0xFF1E1E1E),
      F1Colors.trackBlack,
    ],
    stops: [0.0, 0.25, 0.5, 0.75, 1.0],
  );

  static const podium = LinearGradient(
    colors: [F1Colors.winnersGold, Color(0xFFB8860B)],
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
  );

  static const energy = LinearGradient(
    colors: [F1Colors.speedBlue, F1Colors.speedPurple],
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
  );

  static const neon = LinearGradient(
    colors: [F1Colors.techBlue, F1Colors.speedPurple, F1Colors.energyGreen],
    begin: Alignment.topLeft,
    end: Alignment.bottomRight,
  );

  static const flash = LinearGradient(
    colors: [
      Color(0xFFFFD700),
      Color(0xFFFFF000),
      Color(0xFFFFD700),
      Color(0xFFFFF000),
    ],
    stops: [0.0, 0.3, 0.6, 1.0],
  );
}
