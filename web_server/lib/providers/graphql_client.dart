import 'package:graphql_flutter/graphql_flutter.dart';

class GraphQLClientProvider {
  static final HttpLink httpLink = HttpLink('http://localhost:8080/');

  static GraphQLClient client = GraphQLClient(
    link: httpLink,
    cache: GraphQLCache(),
  );
}
