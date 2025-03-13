import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'screens/F1BettingLanding.dart';
import 'theme/theme.dart';
import 'screens/live_race_screen.dart';
import 'screens/api_test_screen.dart';

void main() async {
  await initHiveForFlutter();
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
      routes: {
        '/live': (context) => const LiveRaceScreen(),
        '/test': (context) => const ApiTestScreen(),
      },
    );
  }
}
