import 'dart:html';
import 'package:polymer/init.dart';
import 'package:polymer/polymer.dart';
import 'package:logging/logging.dart';
import 'package:gae_go_dart_skeleton/components/main_app.dart';

main() async {
  Logger.root.level = Level.INFO;
  Logger.root.onRecord.listen((log) {
    print(log);
  });
  await initPolymer();
}
