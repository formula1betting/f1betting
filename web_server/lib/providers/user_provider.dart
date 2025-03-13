import 'package:graphql_flutter/graphql_flutter.dart';
import './graphql_client.dart';

class UserProvider {
  static Future<Map<String, dynamic>?> getUserById(String id) async {
    const String query = '''
        query GetUser(\$id: ID!) {
          user(id: \$id) {
            id
            fullName
            email
            username
            balance
          }
        }
    ''';

    final QueryResult result = await GraphQLClientProvider.client.query(
      QueryOptions(
        document: gql(query),
        variables: {'id': id},
      ),
    );

    if (result.hasException) {
      throw result.exception!;
    }

    return result.data?['user'];
  }

  static Future<String> createUser(Map<String, dynamic> userInput) async {
    const String mutation = '''
      mutation CreateUser(\$input: UserInput!) {
        createUser(input: \$input)
      }
    ''';

    final QueryResult result = await GraphQLClientProvider.client.mutate(
      MutationOptions(
        document: gql(mutation),
        variables: {'input': userInput},
      ),
    );

    if (result.hasException) {
      throw result.exception!;
    }

    return result.data?['createUser'];
  }
}
