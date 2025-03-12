import 'package:flutter/material.dart';
import '../theme/colors.dart';

class AutoScrollingFeatures extends StatefulWidget {
  final List<Map<String, dynamic>> features;

  const AutoScrollingFeatures({super.key, required this.features});

  @override
  State<AutoScrollingFeatures> createState() => _AutoScrollingFeaturesState();
}

class _AutoScrollingFeaturesState extends State<AutoScrollingFeatures>
    with SingleTickerProviderStateMixin {
  late ScrollController _scrollController;
  late AnimationController _animationController;

  @override
  void initState() {
    super.initState();
    _scrollController = ScrollController();
    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 60),
    )..addListener(_scrollListener);

    WidgetsBinding.instance.addPostFrameCallback((_) {
      _startScrolling();
    });
  }

  void _scrollListener() {
    if (_scrollController.hasClients) {
      double maxScroll = _scrollController.position.maxScrollExtent;
      double currentScroll = _scrollController.position.pixels;

      if (currentScroll >= maxScroll) {
        _scrollController.animateTo(0,
            duration: const Duration(milliseconds: 500),
            curve: Curves.easeInOut);
      } else {
        _scrollController.animateTo(maxScroll * _animationController.value,
            duration: const Duration(milliseconds: 500), curve: Curves.linear);
      }
    }
  }

  void _startScrolling() {
    _animationController.repeat();
  }

  @override
  void dispose() {
    _animationController.dispose();
    _scrollController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 80,
      child: ListView.builder(
        controller: _scrollController,
        scrollDirection: Axis.horizontal,
        padding: const EdgeInsets.symmetric(horizontal: 40),
        itemCount:
            widget.features.length * 3, // Triple the items for smoother looping
        itemBuilder: (context, index) {
          final feature = widget.features[index % widget.features.length];
          return MouseRegion(
            onEnter: (_) => setState(() {}),
            onExit: (_) => setState(() {}),
            child: AnimatedContainer(
              duration: const Duration(milliseconds: 200),
              margin: const EdgeInsets.only(right: 30),
              child: Row(
                children: [
                  TweenAnimationBuilder(
                    tween: ColorTween(
                      begin: F1Colors.racingRed,
                      end: F1Colors.racingRed.withOpacity(0.7),
                    ),
                    duration: const Duration(milliseconds: 500),
                    builder: (context, Color? color, child) {
                      return Icon(
                        feature['icon'] as IconData,
                        color: color,
                        size: 28,
                      );
                    },
                  ),
                  const SizedBox(width: 15),
                  Text(
                    feature['text'] as String,
                    style: const TextStyle(
                      color: Colors.white,
                      fontSize: 18,
                    ),
                  ),
                ],
              ),
            ),
          );
        },
      ),
    );
  }
}
