import 'package:flutter/material.dart';

import '../theme/colors.dart';
import '../theme/gradients.dart';

class BettingCard extends StatefulWidget {
  final String driverName;
  final String odds;

  const BettingCard({
    super.key,
    required this.driverName,
    required this.odds,
  });

  @override
  BettingCardState createState() => BettingCardState();
}

class BettingCardState extends State<BettingCard>
    with SingleTickerProviderStateMixin {
  late final AnimationController _controller;
  late final Animation<double> _glowAnimation;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      duration: const Duration(seconds: 2),
      vsync: this,
    )..repeat(reverse: true);
    _glowAnimation = Tween<double>(begin: 0.5, end: 2.0).animate(_controller);
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Get screen width
    final screenWidth = MediaQuery.of(context).size.width;
    final bool isMobile = screenWidth < 600;

    return isMobile ? _buildMobileCard() : _buildDesktopCard();
  }

  Widget _buildMobileCard() {
    return Container(
      decoration: BoxDecoration(
        color: const Color(0xFF333333),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: F1Colors.racingRed,
          width: 1.5,
        ),
        boxShadow: [
          BoxShadow(
            color: F1Colors.redGlow,
            blurRadius: _glowAnimation.value * 5,
            spreadRadius: _glowAnimation.value,
          ),
        ],
      ),
      padding: const EdgeInsets.all(10),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            widget.driverName,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 14,
              fontWeight: FontWeight.bold,
            ),
            overflow: TextOverflow.ellipsis,
            maxLines: 1,
          ),
          const SizedBox(height: 4),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text(
                'Odds: ',
                style: TextStyle(
                  color: Colors.white70,
                  fontSize: 12,
                ),
              ),
              ShaderMask(
                shaderCallback: (bounds) =>
                    F1Gradients.flash.createShader(bounds),
                child: Text(
                  widget.odds,
                  style: const TextStyle(
                    color: Colors.white,
                    fontSize: 20,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildDesktopCard() {
    return Container(
      decoration: BoxDecoration(
        color: const Color(0xFF333333),
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: F1Colors.racingRed,
          width: 1.5,
        ),
        boxShadow: [
          BoxShadow(
            color: F1Colors.redGlow,
            blurRadius: _glowAnimation.value * 5,
            spreadRadius: _glowAnimation.value,
          ),
        ],
      ),
      padding: const EdgeInsets.all(20),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            widget.driverName,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 18,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 8),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text(
                'Odds: ',
                style: TextStyle(
                  color: Colors.white70,
                  fontSize: 14,
                ),
              ),
              ShaderMask(
                shaderCallback: (bounds) =>
                    F1Gradients.flash.createShader(bounds),
                child: Text(
                  widget.odds,
                  style: const TextStyle(
                    color: Colors.white,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
