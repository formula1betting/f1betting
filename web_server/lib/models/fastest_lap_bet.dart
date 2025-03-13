class FastestLapBet {
  final String id;
  final int driverId;
  final double amount;
  final String status;

  FastestLapBet({
    required this.id,
    required this.driverId,
    required this.amount,
    required this.status,
  });

  factory FastestLapBet.fromJson(Map<String, dynamic> json) {
    return FastestLapBet(
      id: json['id'],
      driverId: json['driverId'],
      amount: json['amount'].toDouble(),
      status: json['status'],
    );
  }
}
