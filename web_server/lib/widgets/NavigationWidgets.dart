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
            _buildDrawerItem('Home', context: context),
            _buildDrawerItem('Live Betting', context: context),
            _buildDrawerItem('Races', context: context),
            _buildDrawerItem('Drivers', context: context),
            _buildDrawerItem('Constructors', context: context),
            _buildDrawerItem('Promotions', context: context),
            _buildDrawerItem('My Account', context: context),
          ],
        ),
      ),
    );
  }

  Widget _buildDrawerItem(String text, {required BuildContext context}) {
    return ListTile(
      title: Text(
        text,
        style: const TextStyle(color: Colors.white),
      ),
      onTap: () {
        // Add navigation based on text
        switch (text) {
          case 'Home':
            Navigator.pushNamed(context, '/');
            break;
          case 'Live Betting':
            Navigator.pushNamed(context, '/live');
            break;
          case 'Races':
            Navigator.pushNamed(context, '/races');
            break;
          case 'Season Betting':
            Navigator.pushNamed(context, '/season');
            break;
          default:
            break;
        }
      },
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
              _buildNavLink('Live Betting', context),
              _buildNavLink('Races', context),
              _buildNavLink('Drivers', context),
              _buildNavLink('Constructors', context),
              _buildNavLink('Promotions', context),
              _buildNavLink('My Account', context),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildNavLink(String text, BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 10),
      child: TextButton(
        onPressed: () {
          // Add navigation based on text
          switch (text) {
            case 'Live Betting':
              Navigator.pushNamed(context, '/live');
              break;
            case 'Races':
              Navigator.pushNamed(context, '/races');
              break;
            case 'Season Betting':
              Navigator.pushNamed(context, '/season');
              break;
            default:
              break;
          }
        },
        style: TextButton.styleFrom(
          foregroundColor: Colors.white,
        ),
        child: Text(text),
      ),
    );
  }
}
