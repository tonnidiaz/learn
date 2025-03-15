import 'dart:convert';

import 'package:tut1/tut1.dart' as tut1;

void main(List<String> arguments) {
  print('Hello world: ${tut1.calculate()}!');
  tut1.startWsClient().then((channel) {
    print("\nSending test msg...");
    final Map<String, dynamic> msg = {
      "event": "hello",
      "data": "Hello dawgie!"
    };
    channel.sink.add(jsonEncode(msg));
  });
}
