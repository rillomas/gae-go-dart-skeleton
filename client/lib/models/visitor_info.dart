part of models;

// Visitor information
class VisitorInfo {
  VisitorInfo(this.count);
  VisitorInfo.fromJson(Map map) {
    count = map[countTag];
  }

  int count;

  dynamic toJson() => {
    countTag: count
  };

  static const String countTag = "Count";
}
