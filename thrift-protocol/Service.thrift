namespace go services
namespace py services
namespace java  sebastians.sportan.networking

//it is critical that structs are defined before services using them!

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

struct UserProfile {
    1: optional string identifier,
    2: optional string username (go.tag = "bson:\"username,omitempty\""),
    3: optional binary profilepicture,

}

struct User {
    1: optional string identifier (go.tag = "bson:\"username,omitempty\""),
    2: optional string password (go.tag = "bson:\"password,omitempty\""),
    3: optional string role (go.tag = "bson:\"role,omitempty\""),
    4: optional UserProfile profile (go.tag = "bson:\"profile,omitempty\""),

}

struct ThriftToken {
    1: required string token (go.tag = "bson:\"token,omitempty\""),
    2: required i64 validityDuration(go.tag = "bson:\"validity,omitempty\""),
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
  1: optional string id,
  2: optional string name,
  3: optional Image icon,
  4: optional string iconid,
}

struct Area {
  1: optional string id,
  2: optional string title,
  3: optional list<string> sports,
  4: optional Coordinate center,
  5: optional list<Coordinate> coords,
  6: optional string description,
  7: optional string cityid,
  8: optional string imageid,

}

struct City {
  1: required string id,
  2: required string name,
  3: optional list<Coordinate> coords,
  4: optional Coordinate center,
}





service SportSvc {
  Sport createSport(1:string token, 2: Sport sport);
  list<Sport> getAllSports(1:string bla);
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

//    UserProfile getProfile(1: string useridentifier);
    void setProfile(1: string token, 2: UserProfile profile);

    ThriftToken requestToken(1: string username, 2: string plain_pw);


    void setAdmin(1: string token, 2: string userid);

}

service AreaSvc {
    //void createArea(1: string title, 2: string description, 3: list<Coordinate> coords);
    void createArea(1: string token, 2:Area area);
    void updateArea(1: string token, 2:Area area);

    void deleteArea(1:string token, 2: Area area);

    list<Area> getNearBy(1: Coordinate coordinate, 2: i32 limit);
    list<Area> getAllAreasInCity(1: string cityid);
    Area getAreaById(1:string id);

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
