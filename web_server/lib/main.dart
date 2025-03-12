import 'package:flutter/material.dart';
import 'screens/F1BettingLanding.dart'; // Updated import path
import 'theme/theme.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'F1 Betting',
      theme: F1Theme.darkTheme,
      home: const F1BettingLanding(),
    );
  }
}
