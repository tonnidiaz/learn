import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

int calculate() {
  return 6 * 7;
}

Future<WebSocketChannel> startWsClient() async {
  final wsUrl = Uri.parse('ws://localhost:8080/ws');
  final channel = WebSocketChannel.connect(wsUrl);

  await channel.ready;
  print("Ws connection established: $wsUrl\n");
  channel.stream.listen((message) {
    print("\n[ws] received: $message");
    // channel.sink.add('received!');
    // channel.sink.close(status.goingAway);
  });

  // Wait for 5 seconds before closing the connection
  // await Future.delayed(Duration(seconds: 5));
  // channel.sink.close(status.goingAway);

  return channel;
}
