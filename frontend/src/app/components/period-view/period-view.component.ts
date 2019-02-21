import { Component, OnInit, ViewChild } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { FormControl } from '@angular/forms';
import { MatTableDataSource, MatSort } from '@angular/material';

import { MatchUpLog } from '../../models/match-up-log';
import * as moment from 'moment';

@Component({
  selector: 'app-components-period-view',
  templateUrl: './period-view.component.html',
  styleUrls: ['./period-view.component.css'],
  entryComponents: []
})
export class AppComponentsPeriodViewComponent implements OnInit {

  constructor(private http: HttpClient) {}

  baseUrl = ''; // 'http://localhost:8080';

  startDate: FormControl = new FormControl(moment().add(-1, 'days'));
  endDate: FormControl = new FormControl(moment());
  winLosss: any[];
  sortedData: any[];

  displayedColumns = ['type', 'win', 'lose', 'percentage'];
  dataSource;
  @ViewChild(MatSort) sort: MatSort;

  ngOnInit(): void {
    this.search();
  }

  search(): void {
    const options = {
      params: new HttpParams()
        .set('startDate', this.startDate.value.format('YYYY-MM-DD'))
        .set('endDate', this.endDate.value.format('YYYY-MM-DD'))
        .set('format', '殿堂')
    };
    this.http.get<MatchUpLog>(this.baseUrl + '/api/v1/matchupLog', options).subscribe((res: any) => {
      const types: string[] = [];
      const winLosss: {[key: string]: any}[] = [];
      res.forEach((m: MatchUpLog) => {
        let deckType = '';
        if (m.PlayerType !== '' && m.PlayerRace !== '') {
          const str = this.isMobile() ? '\n' : '';
          deckType = m.PlayerType + str + '(' + m.PlayerRace + ')';
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
      winLosss.forEach((w: any) => {
        w.games = w.win + w.lose;
        w.percentage = w.win / (w.win + w.lose);
      });
      this.winLosss = winLosss;
      this.sortedData = this.winLosss.slice();
      this.dataSource = new MatTableDataSource(this.winLosss);
      this.dataSource.sort = this.sort;
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

  private isMobile(): boolean {
    const ua = navigator.userAgent.toLowerCase();
    const sdev = /iphone;|(android|nokia|blackberry|bb10;).+mobile|android.+fennec|opera.+mobi|windows phone|symbianos/;
    return sdev.test(ua);
  }
}
