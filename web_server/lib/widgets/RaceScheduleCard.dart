import 'package:flutter/material.dart';

import '../theme/colors.dart';

class RaceScheduleCard extends StatelessWidget {
  const RaceScheduleCard({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(bottom: 10),
      padding: const EdgeInsets.all(20),
      decoration: BoxDecoration(
        color: const Color(0xFF333333),
        borderRadius: BorderRadius.circular(10),
      ),
      child: const Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Text(
            'Monaco Grand Prix',
            style: TextStyle(color: Colors.white, fontSize: 18),
          ),
          Text(
            'May 28, 2024',
            style: TextStyle(color: F1Colors.racingRed, fontSize: 18),
          ),
        ],
      ),
    );
  }
}
