part of models;

// Visitor information
class VisitorInfo extends JsProxy {
  VisitorInfo(this.count);
  VisitorInfo.fromJson(Map map) {
    count = map[countTag];
  }

  @reflectable int count;

  dynamic toJson() => {
    countTag: count
  };

  static const String countTag = "Count";
}
