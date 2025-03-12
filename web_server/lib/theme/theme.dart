import 'package:flutter/material.dart';
import 'colors.dart';
import 'package:google_fonts/google_fonts.dart';

class F1Theme {
  static ThemeData get darkTheme => ThemeData(
        brightness: Brightness.dark,
        primaryColor: F1Colors.racingRed,
        scaffoldBackgroundColor: F1Colors.trackBlack,
        fontFamily: GoogleFonts.oxanium().fontFamily,
        colorScheme: const ColorScheme.dark(
          primary: F1Colors.racingRed,
          secondary: F1Colors.techBlue,
          surface: F1Colors.asphaltGray,
          error: F1Colors.racingRed,
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ElevatedButton.styleFrom(
            backgroundColor: F1Colors.racingRed,
            foregroundColor: F1Colors.checkpointWhite,
            elevation: 4,
          ),
        ),
        textTheme: TextTheme(
          displayLarge: GoogleFonts.oxanium(
            color: F1Colors.textPrimary,
            fontSize: 48,
            fontWeight: FontWeight.bold,
          ),
          displayMedium: GoogleFonts.oxanium(
            color: F1Colors.textPrimary,
            fontSize: 36,
            fontWeight: FontWeight.bold,
          ),
          bodyLarge: GoogleFonts.oxanium(
            color: F1Colors.textPrimary,
            fontSize: 16,
          ),
          bodyMedium: GoogleFonts.oxanium(
            color: F1Colors.textSecondary,
            fontSize: 14,
          ),
        ),
        dividerTheme: DividerThemeData(
          color: F1Colors.divider,
          thickness: 1,
        ),
      );
}
