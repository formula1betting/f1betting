import 'package:flutter/material.dart';
import '../theme/gradients.dart';

class MobileAppBar extends StatelessWidget implements PreferredSizeWidget {
  const MobileAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: Colors.transparent,
      title: const Text('F1Betting'),
      elevation: 0,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}

class MobileDrawer extends StatelessWidget {
  const MobileDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Container(
        color: const Color(0xFF2C2C2C),
        child: ListView(
          children: [
            const DrawerHeader(
              decoration: BoxDecoration(
                gradient: F1Gradients.track,
              ),
              child: Text(
                'F1Betting',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                ),
              ),
            ),
            _buildDrawerItem('Home'),
            _buildDrawerItem('Live Betting'),
            _buildDrawerItem('Races'),
            _buildDrawerItem('Drivers'),
            _buildDrawerItem('Constructors'),
            _buildDrawerItem('Promotions'),
            _buildDrawerItem('My Account'),
          ],
        ),
      ),
    );
  }

  Widget _buildDrawerItem(String text) {
    return ListTile(
      title: Text(
        text,
        style: const TextStyle(color: Colors.white),
      ),
      onTap: () {},
    );
  }
}

class DesktopNavigationBar extends StatelessWidget {
  const DesktopNavigationBar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const Text(
            'F1Betting',
            style: TextStyle(
              color: Colors.white,
              fontSize: 24,
              fontWeight: FontWeight.bold,
            ),
          ),
          Row(
            children: [
              _buildNavLink('Home'),
              _buildNavLink('Live Betting'),
              _buildNavLink('Races'),
              _buildNavLink('Drivers'),
              _buildNavLink('Constructors'),
              _buildNavLink('Promotions'),
              _buildNavLink('My Account'),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildNavLink(String text) {
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 10),
      child: TextButton(
        onPressed: () {},
        style: TextButton.styleFrom(
          foregroundColor: Colors.white,
        ),
        child: Text(text),
      ),
    );
  }
}
