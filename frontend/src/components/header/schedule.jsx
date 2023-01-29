import '../../styles/header/schedule.css';
import boston_logo from '../../pictures/nba_team_logos/boston-celtics-logo.png';
import {useEffect, useState} from "react";
import {getSchedule} from "../../api/schedule_api";

const MatchDateAndTime = ({date, time}) => {

    return (
        <div className="match-date">
            <div className="match-date-day">
                {date}
            </div>
            <div className="match-date-time">
                {time}
            </div>
        </div>
    );
}


const TeamWinsLoses = ({teamWins, teamLoses}) => {
    return (
      <div className="wins-loses">
          {teamWins} - {teamLoses}
      </div>
    );
}

const TeamScore = ({score}) => {
    return (
      <div className="score">
          {score}
      </div>
    );
}

const Team = ({teamName}) => {
    return (
        <div className="team">
            <img className="team-logo"
                src={ boston_logo } alt="Boston Celtics team logo"/>
            <div className="team-name">
                {teamName}
            </div>
        </div>
    );
}

const ScheduledMatch = ({match}) => {
    return (
        <div className="schedule-item">
            <MatchDateAndTime date={match.getMatchDate()} time={match.getMatchTime()}/>
            <div className="matchup">
                <Team teamName={match.homeTeam.acronym}/>
                {match.homeScore ? <TeamScore score={match.homeScore}/> : <TeamWinsLoses teamWins={match.homeTeam.wins} teamLoses={match.homeTeam.losses}/>}
                <Team teamName={match.awayTeam.acronym}/>
                {match.awayScore ? <TeamScore score={match.awayScore}/> : <TeamWinsLoses teamWins={match.awayTeam.wins} teamLoses={match.awayTeam.losses}/>}
            </div>
        </div>
    );
}

const EmptySchedule = () => {
    return (
        <div className="empty-schedule">
            No matches scheduled
        </div>
    );
}


export const Schedule = () => {
    const [schedule, setSchedule] = useState([]);

    useEffect(() => {
        getSchedule().then((r) => setSchedule(r))
    },[])

    return (
        <div className="schedule">
            {
                schedule.length ?
                    schedule.map(
                        (scheduledMatch) =>
                            <ScheduledMatch key={scheduledMatch.id}
                                            match={scheduledMatch}/>)
                : <EmptySchedule />
            }
        </div>
    );
}