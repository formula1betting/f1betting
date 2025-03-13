import 'package:flutter/material.dart';
import '../theme/colors.dart';

class LiveRaceScreen extends StatelessWidget {
  const LiveRaceScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        backgroundColor: Colors.transparent,
        title: Row(
          children: [
            const Text('Live Race'),
            const SizedBox(width: 10),
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
              decoration: BoxDecoration(
                color: F1Colors.racingRed,
                borderRadius: BorderRadius.circular(4),
              ),
              child: const Text('LIVE'),
            ),
          ],
        ),
      ),
      body: Column(
        children: [
          _buildLiveTimerSection(),
          Expanded(
            child: Row(
              children: [
                Expanded(
                  flex: 2,
                  child: _buildLiveLeaderboard(),
                ),
                Expanded(
                  flex: 1,
                  child: _buildLiveBettingOptions(),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildLiveTimerSection() {
    return Container(
      padding: const EdgeInsets.all(16),
      color: F1Colors.asphaltGray,
      child: const Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(
            'Lap 23/58',
            style: TextStyle(color: Colors.white, fontSize: 20),
          ),
          Text(
            '1:23.456',
            style: TextStyle(color: F1Colors.racingRed, fontSize: 20),
          ),
        ],
      ),
    );
  }

  Widget _buildLiveLeaderboard() {
    return ListView.builder(
      itemCount: 20,
      itemBuilder: (context, index) {
        return ListTile(
          leading: Text(
            '${index + 1}',
            style: const TextStyle(color: Colors.white),
          ),
          title: const Text(
            'Driver Name',
            style: TextStyle(color: Colors.white),
          ),
          trailing: const Text(
            '+1.234',
            style: TextStyle(color: F1Colors.racingRed),
          ),
        );
      },
    );
  }

  Widget _buildLiveBettingOptions() {
    return ListView(
      padding: const EdgeInsets.all(8),
      children: [
        _buildLiveBetCard('Next Overtake', '2.50'),
        _buildLiveBetCard('Safety Car', '3.75'),
        _buildLiveBetCard('Fastest Lap', '1.95'),
      ],
    );
  }

  Widget _buildLiveBetCard(String title, String odds) {
    return Card(
      color: F1Colors.racingRed,
      child: Padding(
        padding: const EdgeInsets.all(8),
        child: Column(
          children: [
            Text(
              title,
              style: const TextStyle(color: Colors.white),
            ),
            Text(
              odds,
              style: const TextStyle(
                color: F1Colors.textPrimary,
                fontWeight: FontWeight.bold,
                fontSize: 20,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
