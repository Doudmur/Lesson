let movies = [
    {
        "name": "Step Up 3D",
        "genre": "music",
        "duration": 101
    },
    {
        "name": "The Maze Runner",
        "genre": "fantastic",
        "duration": 113
    },
    {
        "name": "Avengers: Endgame",
        "genre": "action",
        "duration": 181
    }
]
let sum = 0;

for (let i = 0; i < movies.length; i++) {
    sum += movies[i]['duration']
}


avg_duration = sum / movies.length;

let iron_man = {
    "name": "Iron Man",
    "genre": "adventure",
    "duration": 121
}

movies.push(iron_man)


for (let i = 0; i < movies.length; i++) {
    document.write("<h5>Movie: " + movies[i]["name"] + "</h5>");
    document.write("<h5>Genre: " + movies[i]["genre"] + "</h5>");
    document.write("<h5>Duration (minutes): " + movies[i]["duration"] + "");
}

document.write("<h5>Average duration: " + avg_duration + " minutes</h5>");

