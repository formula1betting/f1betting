import 'package:flutter/material.dart';
import '../theme/colors.dart';
import '../theme/gradients.dart';
import '../widgets/BettingCard.dart';
import '../widgets/BettingOptionCard.dart';
import '../widgets/RaceScheduleCard.dart';
import '../widgets/StepCard.dart';
import '../widgets/DriverStandings.dart';
import '../widgets/NavigationWidgets.dart';
import '../widgets/AutoScrollingFeatures.dart';

class F1BettingLanding extends StatelessWidget {
  const F1BettingLanding({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final isMobile = MediaQuery.of(context).size.width < 600;

    return Scaffold(
      appBar: isMobile ? const MobileAppBar() : null,
      drawer: isMobile ? const MobileDrawer() : null,
      body: Stack(
        children: [
          // Background gradient with dark theme
          Container(
            decoration: const BoxDecoration(
              gradient: F1Gradients.track,
            ),
          ),

          // Main content
          SingleChildScrollView(
            child: Column(
              children: [
                if (!isMobile) const DesktopNavigationBar(),

                // Hero section
                SizedBox(
                  height: isMobile ? 400 : 600,
                  child: Stack(
                    children: [
                      // Background racing image with overlay
                      Positioned.fill(
                        child: ShaderMask(
                          shaderCallback: (bounds) => const LinearGradient(
                            begin: Alignment.topCenter,
                            end: Alignment.bottomCenter,
                            colors: [Colors.transparent, Colors.black],
                          ).createShader(bounds),
                          blendMode: BlendMode.darken,
                          child: Image.asset(
                            'assets/f1_hero.jpg',
                            fit: BoxFit.cover,
                          ),
                        ),
                      ),

                      // Hero content
                      Center(
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            Padding(
                              padding:
                                  const EdgeInsets.symmetric(horizontal: 20),
                              child: Text(
                                'FUEL YOUR PASSION WITH EVERY BET',
                                textAlign: TextAlign.center,
                                style: TextStyle(
                                  fontSize: isMobile ? 32 : 48,
                                  fontWeight: FontWeight.bold,
                                  color: Colors.white,
                                ),
                              ),
                            ),
                            const SizedBox(height: 20),
                            Padding(
                              padding:
                                  const EdgeInsets.symmetric(horizontal: 20),
                              child: Text(
                                'Live odds for every Grand Prix, driver, and podium finish',
                                textAlign: TextAlign.center,
                                style: TextStyle(
                                  fontSize: isMobile ? 18 : 24,
                                  color: Colors.white70,
                                ),
                              ),
                            ),
                            const SizedBox(height: 30),
                            ElevatedButton(
                              onPressed: () {},
                              style: ElevatedButton.styleFrom(
                                backgroundColor: const Color(0xFFFF1E1E),
                                padding: const EdgeInsets.symmetric(
                                  horizontal: 40,
                                  vertical: 20,
                                ),
                              ),
                              child: const Text(
                                'JOIN NOW',
                                style: TextStyle(fontSize: 20),
                              ),
                            ),
                            const SizedBox(height: 20),
                            // Current Grand Prix countdown timer
                            _buildCountdownTimer(),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),

                // Key Features
                _buildFeatureSection(),

                // Featured Bets Section
                _buildFeaturedBetsSection(context),

                // Live Betting Highlight
                _buildLiveBettingSection(),

                // Driver/Team Statistics
                _buildDriverTeamStatistics(),

                // Betting Options
                _buildBettingOptions(context),

                // Trust & Security
                _buildTrustSection(),

                // How It Works
                _buildHowItWorks(),

                // Promotions
                _buildPromotions(),

                // Footer
                _buildFooter(),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildCountdownTimer() {
    return TweenAnimationBuilder(
      tween: Tween<double>(begin: 1.0, end: 1.1),
      duration: const Duration(seconds: 1),
      builder: (context, double scale, child) {
        return Transform.scale(
          scale: scale,
          child: Container(
            padding: const EdgeInsets.all(20),
            decoration: BoxDecoration(
              color: const Color(0xFF333333).withAlpha(128),
              borderRadius: BorderRadius.circular(10),
              boxShadow: [
                BoxShadow(
                  color: F1Colors.racingRed.withOpacity(0.3),
                  blurRadius: 20 * scale,
                  spreadRadius: 5,
                ),
              ],
            ),
            child: const Column(
              children: [
                Text(
                  'Next Race Starts In:',
                  style: TextStyle(color: Colors.white, fontSize: 20),
                ),
                Text(
                  '2 days 14 hours',
                  style: TextStyle(
                    color: F1Colors.racingRed,
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    shadows: [
                      Shadow(
                        color: F1Colors.racingRed,
                        blurRadius: 20,
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        );
      },
    );
  }

  Widget _buildFeaturedBetsSection(BuildContext context) {
    final isMobile = MediaQuery.of(context).size.width < 600;

    return Container(
      padding: EdgeInsets.all(isMobile ? 20 : 40),
      child: Column(
        children: [
          Text(
            'THIS WEEKEND',
            style: TextStyle(
              color: Colors.white,
              fontSize: isMobile ? 24 : 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 20),
          const RaceScheduleCard(),
          const SizedBox(height: 20),
          GridView.count(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            crossAxisCount: isMobile ? 3 : 5,
            childAspectRatio: 1.5,
            mainAxisSpacing: 10,
            crossAxisSpacing: 10,
            children: const [
              BettingCard(driverName: 'Max Verstappen', odds: '1.50'),
              BettingCard(driverName: 'Lewis Hamilton', odds: '2.25'),
              BettingCard(driverName: 'Charles Leclerc', odds: '3.75'),
              BettingCard(driverName: 'Lando Norris', odds: '4.50'),
              BettingCard(driverName: 'George Russell', odds: '5.00'),
              BettingCard(driverName: 'Carlos Sainz', odds: '5.50'),
              BettingCard(driverName: 'Fernando Alonso', odds: '6.00'),
              BettingCard(driverName: 'Oscar Piastri', odds: '6.50'),
              BettingCard(driverName: 'Sergio Perez', odds: '7.00'),
              BettingCard(driverName: 'Pierre Gasly', odds: '8.00'),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildLiveBettingSection() {
    return Container(
      padding: const EdgeInsets.all(40),
      color: const Color(0xFF2C2C2C),
      child: const Column(
        children: [
          Text(
            'BETTING LIVE NOW',
            style: TextStyle(
              color: Colors.white,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          SizedBox(height: 20),
          Text(
            'In-race betting odds and updates',
            style: TextStyle(color: Colors.white70, fontSize: 18),
          ),
          SizedBox(height: 20),
          // Add animated graphics and quick-access buttons here
        ],
      ),
    );
  }

  Widget _buildDriverTeamStatistics() {
    return Container(
      padding: const EdgeInsets.all(40),
      child: const Column(
        children: [
          Text(
            'DRIVER STANDINGS',
            style: TextStyle(
              color: Colors.white,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          SizedBox(height: 20),
          DriverStandings(),
        ],
      ),
    );
  }

  Widget _buildFeatureSection() {
    final features = [
      {'icon': Icons.speed, 'text': 'Live In-Race Betting'},
      {'icon': Icons.psychology, 'text': 'AI-Powered Odds'},
      {'icon': Icons.payments, 'text': 'Instant Payouts'},
      {'icon': Icons.diamond, 'text': 'Exclusive VIP Perks'},
      {'icon': Icons.timer, 'text': 'Real-time Updates'},
      {'icon': Icons.trending_up, 'text': 'Dynamic Odds'},
      {'icon': Icons.track_changes, 'text': 'Track Conditions'},
      {'icon': Icons.insights, 'text': 'Performance Analytics'},
      {'icon': Icons.price_check, 'text': 'Best Odds Guaranteed'},
      {'icon': Icons.military_tech, 'text': 'VIP Rewards Program'},
      {'icon': Icons.notifications_active, 'text': 'Race Alerts'},
      {'icon': Icons.history, 'text': 'Historical Stats'},
    ];

    return Container(
      padding: const EdgeInsets.symmetric(vertical: 40),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Padding(
            padding: EdgeInsets.symmetric(horizontal: 40),
            child: Text(
              'FEATURES',
              style: TextStyle(
                color: Colors.white,
                fontSize: 32,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          const SizedBox(height: 20),
          AutoScrollingFeatures(features: features),
        ],
      ),
    );
  }

  Widget _buildBettingOptions(BuildContext context) {
    final isMobile = MediaQuery.of(context).size.width < 600;

    return Container(
      padding: EdgeInsets.all(isMobile ? 20 : 40),
      color: const Color(0xFF2C2C2C),
      child: Column(
        children: [
          Text(
            'BETTING OPTIONS',
            style: TextStyle(
              color: Colors.white,
              fontSize: isMobile ? 24 : 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 40),
          GridView.count(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            crossAxisCount: isMobile ? 2 : 5, // Changed from 2 to 4
            childAspectRatio:
                isMobile ? 1.0 : 1.2, // Adjusted for better proportions
            mainAxisSpacing: 10, // Reduced spacing
            crossAxisSpacing: 10, // Reduced spacing
            children: const [
              BettingOptionCard(
                title: 'Race Winners',
                description:
                    'Back your favorite driver to take the checkered flag',
              ),
              BettingOptionCard(
                title: 'Podium Finishes',
                description: 'Predict who\'ll make it to the top three',
              ),
              BettingOptionCard(
                title: 'Qualifying Leaders',
                description: 'Bet on pole position and qualifying results',
              ),
              BettingOptionCard(
                title: 'Season Champions',
                description: 'Long-term bets on championship winners',
              ),
              BettingOptionCard(
                title: 'Fastest Lap',
                description:
                    'Predict which driver will set the fastest lap time',
              ),
              BettingOptionCard(
                title: 'Safety Car',
                description: 'Bet on safety car appearances during the race',
              ),
              BettingOptionCard(
                title: 'First Pit Stop',
                description: 'Predict the first driver to make a pit stop',
              ),
              BettingOptionCard(
                title: 'Head-to-Head',
                description: 'Bet on direct driver matchups and performance',
              ),
              BettingOptionCard(
                title: 'Constructor Points',
                description: 'Bet on team performance and points finish',
              ),
              BettingOptionCard(
                title: 'DNF Predictions',
                description: 'Predict which drivers won\'t finish the race',
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildTrustSection() {
    return Container(
      padding: const EdgeInsets.all(40),
      child: Column(
        children: [
          const Text(
            'TRUSTED BY F1 FANS',
            style: TextStyle(
              color: Colors.white,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          const SizedBox(height: 30),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              _buildTrustBadge('UKGC Licensed'),
              _buildTrustBadge('SSL Secured'),
              _buildTrustBadge('McLaren Partner'),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildTrustBadge(String text) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
      decoration: BoxDecoration(
        border: Border.all(color: F1Colors.racingRed),
        borderRadius: BorderRadius.circular(5),
      ),
      child: Text(
        text,
        style: const TextStyle(color: Colors.white),
      ),
    );
  }

  Widget _buildHowItWorks() {
    return Container(
      padding: const EdgeInsets.all(40),
      child: const Column(
        children: [
          Text(
            'HOW IT WORKS',
            style: TextStyle(
              color: Colors.white,
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          SizedBox(height: 30),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              StepCard(step: '1', title: 'Sign Up'),
              StepCard(step: '2', title: 'Deposit & Claim Bonus'),
              StepCard(step: '3', title: 'Place Your Bet'),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildPromotions() {
    return Container(
      padding: const EdgeInsets.all(40),
      child: const Column(
        children: [
          Text(
            'Welcome Bonus: 100% Match up to €100',
            style: TextStyle(
              color: Color(0xFFFF1E1E),
              fontSize: 32,
              fontWeight: FontWeight.bold,
            ),
          ),
          SizedBox(height: 20),
          Text(
            'Race-specific promotional offers',
            style: TextStyle(color: Colors.white, fontSize: 20),
          ),
          // Add race-specific promotional offers here
        ],
      ),
    );
  }

  Widget _buildFooter() {
    return Container(
      padding: const EdgeInsets.all(40),
      color: const Color(0xFF1A1A1A),
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              _buildFooterLink('FAQs'),
              _buildFooterLink('Live Stats'),
              _buildFooterLink('Affiliate Program'),
            ],
          ),
          const SizedBox(height: 20),
          const Text(
            'Bet responsibly. 18+ only.',
            style: TextStyle(color: Colors.white70),
          ),
        ],
      ),
    );
  }

  Widget _buildFooterLink(String text) {
    return Text(
      text,
      style: const TextStyle(
        color: Colors.white,
        decoration: TextDecoration.underline,
      ),
    );
  }
}
