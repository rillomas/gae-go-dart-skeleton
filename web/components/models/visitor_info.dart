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

	static const string countTag = "Count";
}
