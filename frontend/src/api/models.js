export class Match {
    constructor(data) {
        this.id = data.id;
        this.date = new Date(data.date);
        this.homeTeam = data.home_team;
        this.homeScore = data.home_score;
        this.awayTeam = data.away_team;
        this.awayScore = data.away_score;
    }

    getMatchDate() {
        const day = this.date.getDate().toString().padStart(2, "0");
        const month = (this.date.getMonth() + 1).toString().padStart(2, "0");
        const year = this.date.getFullYear().toString();

        return `${day}.${month}.${year}`
    }
    getMatchTime(){
        const hour = this.date.getHours().toString().padStart(2, "0");
        const minute = this.date.getMinutes().toString().padStart(2, "0");

        return `${hour}:${minute}`
    }
}

