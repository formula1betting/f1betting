import 'package:flutter/material.dart';
import '../providers/user_provider.dart';
import '../providers/betting_provider.dart';

class ApiTestScreen extends StatefulWidget {
  const ApiTestScreen({Key? key}) : super(key: key);

  @override
  State<ApiTestScreen> createState() => _ApiTestScreenState();
}

class _ApiTestScreenState extends State<ApiTestScreen> {
  String _resultText = '';
  bool _isLoading = false;

  Future<void> _testGetUser() async {
    setState(() {
      _isLoading = true;
      _resultText = 'Fetching user...';
    });

    try {
      final user = await UserProvider.getUserById('1');
      setState(() {
        _resultText = 'User Data: ${user.toString()}';
      });
    } catch (e) {
      setState(() {
        _resultText = 'Error: ${e.toString()}';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _testCreateFastestLapBet() async {
    setState(() {
      _isLoading = true;
      _resultText = 'Creating fastest lap bet...';
    });

    try {
      final betInput = {
        'sessionId': 1,
        'driverId': 44,
        'bettingPool': 100,
      };

      final result =
          await BettingProvider.createFastestLapBet('test-user-id', betInput);
      setState(() {
        _resultText = 'Bet created with ID: $result';
      });
    } catch (e) {
      setState(() {
        _resultText = 'Error: ${e.toString()}';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _testGetFastestLapBets() async {
    setState(() {
      _isLoading = true;
      _resultText = 'Fetching fastest lap bets...';
    });

    try {
      final betsData = await BettingProvider.getFastestLapBetsAndPayout('1', 1);
      setState(() {
        _resultText = 'Bets Data: ${betsData.toString()}';
      });
    } catch (e) {
      setState(() {
        _resultText = 'Error: ${e.toString()}';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('API Test Screen'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            ElevatedButton(
              onPressed: _isLoading ? null : _testGetUser,
              child: const Text('Test Get User'),
            ),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: _isLoading ? null : _testCreateFastestLapBet,
              child: const Text('Test Create Fastest Lap Bet'),
            ),
            const SizedBox(height: 8),
            ElevatedButton(
              onPressed: _isLoading ? null : _testGetFastestLapBets,
              child: const Text('Test Get Fastest Lap Bets'),
            ),
            const SizedBox(height: 16),
            if (_isLoading)
              const Center(child: CircularProgressIndicator())
            else
              Expanded(
                child: SingleChildScrollView(
                  child: Text(
                    _resultText,
                    style: const TextStyle(fontSize: 16),
                  ),
                ),
              ),
          ],
        ),
      ),
    );
  }
}
