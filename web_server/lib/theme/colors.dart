import 'package:flutter/material.dart';

class F1Colors {
  // Primary Colors
  static const racingRed = Color(0xFFFF0600); // Brighter red
  static const trackBlack = Color(0xFF080808); // Darker black
  static const asphaltGray = Color(0xFF1A1A1A);
  static const checkpointWhite = Color(0xFFFFFFFF);

  // Secondary Colors
  static const titaniumSilver = Color(0xFFF0F0F0); // Brighter silver
  static const carbonBlack = Color(0xFF202020);
  static const techBlue = Color(0xFF00FFFF); // Neon cyan
  static const energyGreen = Color(0xFF00FF00); // Pure neon green

  // Accent Colors
  static const boostYellow = Color(0xFFFFD100); // Brighter yellow
  static const speedPurple = Color(0xFFFF00FF); // Neon magenta
  static const adrenalineOrange = Color(0xFFFF5500); // Vibrant orange
  static const speedBlue = Color(0xFF00B3FF); // Electric blue
  static const winnersGold = Color(0xFFFFD700);

  // Text Colors
  static const textPrimary = checkpointWhite;
  static const textSecondary = titaniumSilver;
  static final textMeta = titaniumSilver.withOpacity(0.7);

  // Border Colors
  static const borderPrimary = racingRed;
  static final borderSecondary = titaniumSilver.withOpacity(0.2);
  static final divider = titaniumSilver.withOpacity(0.1);

  // Status Colors
  static const success = energyGreen;
  static const warning = boostYellow;
  static const error = racingRed;
  static const inactive = Color(0xFF4A4A4A); // Muted gray

  // Glow Colors
  static final neonGlow = techBlue.withOpacity(0.5);
  static final redGlow = racingRed.withOpacity(0.5);
  static final goldGlow = winnersGold.withOpacity(0.5);
}
