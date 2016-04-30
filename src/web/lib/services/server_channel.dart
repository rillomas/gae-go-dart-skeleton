part of services;

/**
 * Interface class to the server
 */
class ServerChannel {

  static Future<VisitorInfo> getVisitorInfo(String domain) async {
    var url = "${domain}${visitorInfoApiPath}";
    Logger.root.info("getting from ${url}");
    var r = await HttpRequest.request(url);
    if (!(r.readyState == HttpRequest.DONE && (r.status == 200))) {
      Logger.root.warning("Request failed by ${r.status}");
      return null;
    }
    var res = r.responseText;
    Logger.root.info("response: $res");
    int count = -1;
    if (res.isEmpty) {
      Logger.root.warning("received empty string");
      return null;
    }
    Map map = JSON.decode(res);
    var info = new VisitorInfo.fromJson(map);
    return info;
  }

  static Future<HttpRequest> sendVisitorInfo(VisitorInfo info, String domain) {
    var url = "${domain}${visitorInfoApiPath}";
    Logger.root.info("Sending info to ${url}");
    var str = JSON.encode(info);
    return HttpRequest.request(url, method:"POST", sendData: str);
  }

  // Generate the root domain from location information
  static String generateRootDomain(Location location) {
    return "${location.protocol}//${location.host}/";
  }

  static const String visitorInfoApiPath = "api/1/visitorInfo";
}


