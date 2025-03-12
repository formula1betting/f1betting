import 'package:flutter/material.dart';

class BettingOptionCard extends StatelessWidget {
  final String title;
  final String description;

  const BettingOptionCard({
    super.key,
    required this.title,
    required this.description,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(15), // Reduced padding
      decoration: BoxDecoration(
        color: const Color(0xFF333333),
        borderRadius: BorderRadius.circular(8), // Smaller border radius
      ),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            title,
            textAlign: TextAlign.center,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 18, // Smaller font size
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 8), // Reduced spacing
          Text(
            description,
            textAlign: TextAlign.center,
            style: const TextStyle(
              color: Colors.white70,
              fontSize: 12, // Smaller font size
            ),
          ),
        ],
      ),
    );
  }
}
