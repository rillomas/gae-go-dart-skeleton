@HtmlImport('main_app.html')
library main_app;

import 'dart:html';
import 'dart:convert' show JSON;
import 'package:polymer/polymer.dart';
import 'package:logging/logging.dart';
import 'package:web_components/web_components.dart' show HtmlImport;
import 'package:gae_go_dart_skeleton/services/services.dart';
import 'package:gae_go_dart_skeleton/models/models.dart';

@PolymerRegister("main-app")
class MainApp extends PolymerElement {
  MainApp.created(): super.created();

  @property String visitorInfoString;
  @property VisitorInfo visitorInfo;
  @Property(computed: "hasVisitorInfo(visitorInfo)")
  bool hasInfo = false;
  @reflectable hasVisitorInfo(VisitorInfo info) => info != null;

  ready() async {
    var domain = ServerChannel.generateRootDomain(window.location);
    try {
      var map = JSON.decode(visitorInfoString);
      var info = new VisitorInfo.fromJson(map);
      info.count++;
      set("visitorInfo", info);
      // update the visitor info on the server side
      ServerChannel.sendVisitorInfo(info, domain);
    } catch(e) {
      Logger.root.warning("$e");
    }
  }
}
