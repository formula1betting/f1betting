import 'package:flutter/material.dart';
import '../theme/colors.dart';

class DriverStandings extends StatelessWidget {
  const DriverStandings({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: const Color(0xFF333333),
        borderRadius: BorderRadius.circular(10),
      ),
      child: const Column(
        children: [
          DriverStandingRow(
            position: 'P',
            driver: 'DRIVER',
            team: 'TEAM',
            points: 'PTS',
            wins: 'W',
            podiums: 'POD',
            isHeader: true,
          ),
          DriverStandingRow(
            position: '1',
            driver: 'Max Verstappen',
            team: 'Red Bull Racing',
            points: '575',
            wins: '19',
            podiums: '21',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '2',
            driver: 'Sergio Perez',
            team: 'Red Bull Racing',
            points: '285',
            wins: '2',
            podiums: '8',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '3',
            driver: 'Lewis Hamilton',
            team: 'Mercedes',
            points: '234',
            wins: '0',
            podiums: '6',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '4',
            driver: 'Fernando Alonso',
            team: 'Aston Martin',
            points: '206',
            wins: '0',
            podiums: '8',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '5',
            driver: 'Charles Leclerc',
            team: 'Ferrari',
            points: '206',
            wins: '0',
            podiums: '6',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '6',
            driver: 'Lando Norris',
            team: 'McLaren',
            points: '205',
            wins: '0',
            podiums: '7',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '7',
            driver: 'Carlos Sainz',
            team: 'Ferrari',
            points: '200',
            wins: '1',
            podiums: '3',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '8',
            driver: 'George Russell',
            team: 'Mercedes',
            points: '175',
            wins: '0',
            podiums: '2',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '9',
            driver: 'Oscar Piastri',
            team: 'McLaren',
            points: '97',
            wins: '0',
            podiums: '2',
            isHeader: false,
          ),
          DriverStandingRow(
            position: '10',
            driver: 'Lance Stroll',
            team: 'Aston Martin',
            points: '73',
            wins: '0',
            podiums: '0',
            isHeader: false,
          ),
        ],
      ),
    );
  }
}

class DriverStandingRow extends StatelessWidget {
  final String position;
  final String driver;
  final String team;
  final String points;
  final String wins;
  final String podiums;
  final bool isHeader;

  const DriverStandingRow({
    super.key,
    required this.position,
    required this.driver,
    required this.team,
    required this.points,
    required this.wins,
    required this.podiums,
    required this.isHeader,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 12),
      decoration: BoxDecoration(
        border: Border(
          bottom: BorderSide(
            color: F1Colors.borderSecondary,
            width: 1,
          ),
        ),
      ),
      child: Row(
        children: [
          SizedBox(
            width: 30,
            child: Text(
              position,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.textPrimary,
                fontWeight: isHeader ? FontWeight.normal : FontWeight.bold,
              ),
            ),
          ),
          Expanded(
            flex: 3,
            child: Text(
              driver,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.textPrimary,
                fontWeight: isHeader ? FontWeight.normal : FontWeight.bold,
              ),
            ),
          ),
          Expanded(
            flex: 2,
            child: Text(
              team,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.textMeta,
              ),
            ),
          ),
          SizedBox(
            width: 50,
            child: Text(
              points,
              textAlign: TextAlign.right,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.racingRed,
                fontWeight: isHeader ? FontWeight.normal : FontWeight.bold,
              ),
            ),
          ),
          SizedBox(
            width: 40,
            child: Text(
              wins,
              textAlign: TextAlign.right,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.winnersGold,
              ),
            ),
          ),
          SizedBox(
            width: 40,
            child: Text(
              podiums,
              textAlign: TextAlign.right,
              style: TextStyle(
                color: isHeader ? F1Colors.textSecondary : F1Colors.speedBlue,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
