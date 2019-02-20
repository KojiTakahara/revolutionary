import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { MatchUpLog } from './models/match-up-log';
import { FormControl } from '@angular/forms';
import moment = require('moment');

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  constructor(private http: HttpClient) {}

  baseUrl = 'http://localhost:8080';

  startDate: FormControl = new FormControl(moment().add(-30, 'days'));
  endDate: FormControl = new FormControl(moment());
  winLosss: any[];

  ngOnInit(): void {
    this.search();
  }

  search(): void {
    const options = {
      params: new HttpParams()
        .set('startDate', this.startDate.value.format('YYYY-MM-DD'))
        .set('endDate', this.endDate.value.format('YYYY-MM-DD'))
    };
    this.http.get<MatchUpLog>(this.baseUrl + '/api/v1/matchupLog', options).subscribe((res: any) => {
      const types: string[] = [];
      const winLosss: {[key: string]: any}[] = [];
      res.forEach((m: MatchUpLog) => {
        let deckType = '';
        if (m.PlayerType !== '' && m.PlayerRace !== '') {
          deckType = m.PlayerType + '(' + m.PlayerRace + ')';
        } else if (m.PlayerType !== '') {
          deckType = m.PlayerType;
        } else if (m.PlayerRace !== '') {
          deckType = '(' + m.PlayerRace + ')';
        } else {
          deckType = 'その他';
        }
        if (!types.includes(deckType)) {
          types.push(deckType);
          const w = {type: deckType, win: 0, lose: 0};
          this.addCount(w, m);
          winLosss.push(w);
        } else if (deckType !== '') {
          const winLoss = winLosss.filter((w: {[key: string]: any}) => {
            return w.type === deckType;
          })[0];
          this.addCount(winLoss, m);
        }
      });
      this.winLosss = winLosss;
      console.log(types);
    });
  }

  private addCount(winLoss: {[key: string]: any}, m: MatchUpLog) {
    if (!winLoss) {
      return;
    }
    if (m.Win) {
      winLoss.win++;
    } else {
      winLoss.lose++;
    }
  }

}
