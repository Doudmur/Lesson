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
function avg(movies){
    let sum = 0;
    for (let i = 0; i < movies.length; i++) {
        sum += movies[i]['duration'];
    }
    return sum / movies.length;
}

let iron_man = {
    "name": "Iron Man",
    "genre": "adventure",
    "duration": 121
}

movies.push(iron_man)


function write(i, movies){
    return movies[i];
}

for (let i = 0; i < movies.length; i++) {
    let obj = write(i, movies);
    document.write("<h5>Movie: " + obj["name"] + "</h5>");
    document.write("<h5>Genre: " + obj["genre"] + "</h5>");
    document.write("<h5>Duration (minutes): " + obj["duration"] + "");
    document.writeln("")
}


document.write("<h5>Average duration: " + avg(movies) + " minutes</h5>")

