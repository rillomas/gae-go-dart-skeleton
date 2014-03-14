part of services;

/**
 * A class to wrap Async get requests
 */
class AsyncGet<Output> {
  Future<Output> request(String url, Output process(String response)) {
    HttpRequest.request(url).then((r) {
      if (r.readyState == HttpRequest.DONE && (r.status == 200)) {
        Output out = process(r.responseText); // convert the raw response text
        _completer.complete(out);
      }
    });
    return _completer.future;
  }

  Completer<Output> _completer = new Completer<Output>();
}

/**
 * Interface class to the server
 */
class ServerChannel {

  static Future<VisitorInfo> getVisitorInfo(String domain) {
    var url = "${domain}${visitorInfoApiPath}";
    Logger.root.info("getting from ${url}");
    var task = new AsyncGet<VisitorInfo>();
    var f = task.request(url, (res) {
      Logger.root.info("response: $res");
      int count = -1;
      if (res.isEmpty) {
        Logger.root.warning("received empty string");
        return count;
      }
      Map map = JSON.decode(res);
      var info = new VisitorInfo.fromJson(map);
      return info;
    });
    return f;
  }

  static void sendVisitorInfo(VisitorInfo info, String domain) {
    var request = new HttpRequest();
    var url = "${domain}${visitorInfoApiPath}";
    Logger.root.info("Sending info to ${url}");
    request.open("POST", url);
    var str = JSON.encode(info);
    request.send(str);
  }

  // Generate the root domain from location information
  static String generateRootDomain(Location location) {
    return "${location.protocol}//${location.host}/";
  }

  static const String visitorInfoApiPath = "api/1/visitorInfo";
}


