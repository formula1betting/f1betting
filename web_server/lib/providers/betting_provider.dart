import 'package:graphql_flutter/graphql_flutter.dart';
import './graphql_client.dart';

class BettingProvider {
  static Future<String> createFastestLapBet(
      String userId, Map<String, dynamic> betInput) async {
    const String mutation = '''
      mutation CreateFastestLapBet(\$userId: ID!, \$input: FastestLapBetInput!) {
        createFastestLapBet(userId: \$userId, input: \$input)
      }
    ''';

    final QueryResult result = await GraphQLClientProvider.client.mutate(
      MutationOptions(
        document: gql(mutation),
        variables: {
          'userId': userId,
          'input': betInput,
        },
      ),
    );

    if (result.hasException) {
      throw result.exception!;
    }

    return result.data?['createFastestLapBet'];
  }

  static Future<Map<String, dynamic>> getFastestLapBetsAndPayout(
      String userId, int sessionId) async {
    const String query = '''
      query GetFastestLapBets(\$sessionId: Int!, \$userId: ID!) {
        fastestLapBetsAndVisualizedPayout(sessionId: \$sessionId, userId: \$userId) {
          fastestLapBets {
            id
            driverId
            amount
            status
          }
          visualizedPayout {
            driverId
            payout
          }
        }
      }
    ''';

    final QueryResult result = await GraphQLClientProvider.client.query(
      QueryOptions(
        document: gql(query),
        variables: {
          'sessionId': sessionId,
          'userId': userId,
        },
      ),
    );

    if (result.hasException) {
      throw result.exception!;
    }

    return result.data?['fastestLapBetsAndVisualizedPayout'];
  }
}
