namespace go services
namespace py services
namespace java  sebastians.sportan.networking

//it is critical that structs are defined before services using them!

const string ROLE_ADMIN = "admin"
const string ROLE_SUPERADMIN = "superadmin"

exception InvalidOperation {
  1: i32 what,
  2: string why,
}

exception InvalidToken {
  1: string message,
}

struct UserCredentials {
  1: required string identifier,
  2: optional string passwordhash,
}


struct Profile {
    1: optional string identifier (go.tag = "bson:\"identifier,omitempty\""),
    2: optional string username (go.tag = "bson:\"username,omitempty\""),
    3: optional binary profilepicture (go.tag = "bson:\"image_id,omitempty\""),
    4: optional string city_id (go.tag = "bson:\"city_id,omitempty\""),

}

struct Token {
    1: required string token (go.tag = "bson:\"token,omitempty\""),
    2: required i64 validity(go.tag = "bson:\"validity,omitempty\""),
}

struct User {
    1: optional string identifier (go.tag = "bson:\"username,omitempty\""),
    2: optional string password (go.tag = "bson:\"password,omitempty\""),
    3: optional string role (go.tag = "bson:\"role,omitempty\""),
    4: optional Profile profile (go.tag = "bson:\"profile,omitempty\""),
    5: optional Token token (go.tag = "bson:\"token,omitempty\""),
    6: optional list<string> friends (go.tag = "bson:\"friends,omitempty\""),
    7: optional list<string> friendrequests (go.tag = "bson:\"friendrequests,omitempty\""),

}



struct Coordinate {
    1: required double lat,
    2: required double lon,
}

struct Image {
  1: optional string id,
  2: optional string content,
  3: optional string creator,
  4: optional binary bcontent,
}


struct Sport {
  1: optional string id (go.tag = "bson:\"id,omitempty\""),
  2: optional string name (go.tag = "bson:\"name,omitempty\""),
  3: optional Image icon (go.tag = "bson:\"image,omitempty\""),
  4: optional string iconid (go.tag = "bson:\"imageid,omitempty\""),
}

struct Area {
  1: optional string id (go.tag = "bson:\"areaid,omitempty\""),
  2: optional string title (go.tag = "bson:\"title,omitempty\""),
  3: optional list<string> sports (go.tag = "bson:\"sports,omitempty\""),
  4: optional list<double> center (go.tag = "bson:\"center,omitempty\""),
  5: optional list<Coordinate> coords (go.tag = "bson:\"coords,omitempty\""),
  6: optional string description (go.tag = "bson:\"description,omitempty\""),
  7: optional string cityid (go.tag = "bson:\"cityid,omitempty\""),
  8: optional string imageid (go.tag = "bson:\"iamgeid,omitempty\""),

}

struct City {
  1: required string id (go.tag = "bson:\"cityid,omitempty\""),
  2: required string name (go.tag = "bson:\"name,omitempty\""),
  3: optional list<Coordinate> coords (go.tag = "bson:\"coords,omitempty\""),
  4: optional Coordinate center (go.tag = "bson:\"center,omitempty\""),
}


service SportSvc {
  Sport createSport(1:string token, 2: Sport sport);
  list<Sport> getAllSports(1:string bla);
  Sport getSportById(1: string token, 2: string sportid);
  //Sport getSportB
}

service ImageSvc {
  //Image createImage(1: string token, 2: Image image);
  Image getImageById(1: string id);
  Image getThumbnailByImageId(1: string id);
  string createImage(1: string token, 2: Image image);
}



/**
* Userrelated stuff (creation etc)
**/
service UserSvc {
    UserCredentials createUser(1: string password) throws (1:InvalidOperation ouch);

    User getMe(1: string token);
    User getUserById(1: string token, 2: string userid);
    list<User> getFriends(1: string token);
    list<User> getFriendRequests(1: string token);
    void acceptFriendRequest(1: string token, 2: string userid);
    void declineFriendRequest(1: string token, 2: string userid);
    void sendFriendRequest(1: string token, 2: string userid);
    void setProfile(1: string token, 2: Profile profile);


    Token requestToken(1: string username, 2: string plain_pw);
    void setAdmin(1: string token, 2: string userid, 3: bool admin);


}

service AreaSvc {
    //void createArea(1: string title, 2: string description, 3: list<Coordinate> coords);
    void createArea(1: string token, 2:Area area);
    void updateArea(1: string token, 2:Area area);
    void deleteArea(1:string token, 2: Area area);
    Area getAreaById(1:string token, 2: string areaid);

    list<string> getNearBy(1:string token, 2: Coordinate coordinate, 3: i32 limit);

    list<string> getAllAreasInCity(1: string cityid);


    //functions for easy retrieval of new areas!
    //i32 countAreasInCity(1: string cityid);
    //list<Area> getBatchAreasInCity(1: string cityid, i32 offset, i32 limit);

}

service CitySvc {
  void createCity(1: string token, 2: string title, 3: list<Coordinate> coords);
  list<City> getNearBy(1: Coordinate coordinate, 2: i32 limit);
  list<City> getAllCities();
  //City getCityById(1: string id);
}



/*
struct ThriftToken {
    1: required string token,
    2: required i64 validityDuration
}

struct ThriftTask {
    1: required string title,
    2: required string description,
    3: required i32 id,
    4: required i64 duration,
    5: optional list<binary> images
}

service ChallengeSvc {
    string createChallenge(1: string usertoken, 2: string title),
    void addTasks(1: string usertoken, 2: string challengeid, 3: list<ThriftTask> tasks),
    list<string> getMyChallenges(1: string usertoken)

}
*/
