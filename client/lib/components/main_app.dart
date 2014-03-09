library main_app;

import 'dart:html';
import 'package:polymer/polymer.dart';
import 'package:logging/logging.dart';
import 'package:gae_go_dart_skeleton/services/services.dart';
import 'package:gae_go_dart_skeleton/models/models.dart';

@CustomTag("main-app")
class MainApp extends PolymerElement {
	MainApp.created() : super.created() {
		// Setup Logger
		Logger.root.level = Level.INFO;
		Logger.root.onRecord.listen((log) {
			print(log);
		});

		visitorInfo = new VisitorInfo(0);
	}

	@observable
	VisitorInfo visitorInfo;

	@override
	void enteredView() {
		super.enteredView();

		var domain = ServerChannel.generateRootDomain(window.location);
		ServerChannel.getVisitorInfo(domain).then((info) {
			info.count++;
			visitorInfo = info;

			// update the visitor info on the server side
			ServerChannel.sendVisitorInfo(info, domain);
		});

	}
}
