import { Component, EventEmitter, OnInit, Output, ViewChild } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { FormControl } from '@angular/forms';
import { MatTableDataSource, MatSort } from '@angular/material';

import * as moment from 'moment';
import { MatchUpLog } from '../../models';
import { common } from '../../common';

@Component({
  selector: 'app-components-period-view',
  templateUrl: './period-view.component.html',
  styleUrls: ['./period-view.component.css'],
  entryComponents: []
})
export class AppComponentsPeriodViewComponent implements OnInit {

  constructor(private http: HttpClient) {}

  baseUrl = 'http://localhost:8080';

  startDate: FormControl = new FormControl(moment().add(-7, 'days'));
  endDate: FormControl = new FormControl(moment());
  minDate = moment('2019-02-01');
  winLosss: any[];
  sortedData: any[];
  matchUpLogs: MatchUpLog[];
  total = 0;
  format = '殿堂';

  displayedColumns = ['type', 'users', 'win', 'lose', 'percentage'];
  dataSource;

  @Output() moveDetail: EventEmitter<any> = new EventEmitter();
  @ViewChild(MatSort) sort: MatSort;

  ngOnInit(): void {
    this.search();
  }

  search(): void {
    this.total = 0;
    const options = {
      params: new HttpParams()
        .set('startDate', this.startDate.value.format('YYYY-MM-DD'))
        .set('endDate', this.endDate.value.format('YYYY-MM-DD'))
        .set('format', this.format)
    };
    this.http.get<MatchUpLog>(this.baseUrl + '/api/v1/matchupLog', options).subscribe((res: any) => {
      const types: string[] = [];
      const winLosss: {[key: string]: any}[] = [];
      this.matchUpLogs = res;
      res.forEach((m: MatchUpLog) => {
        const deckType = this.getPlayerType(m) || this.getDeckType(m);
        if (!types.includes(deckType)) {
          types.push(deckType);
          const w = {
            type: deckType,
            win: 0,
            lose: 0,
            race: m.PlayerRace,
            tempType: this.getPlayerType(m)
          };
          common.addCount(w, m);
          winLosss.push(w);
        } else if (deckType !== '') {
          const winLoss = winLosss.filter((w: {[key: string]: any}) => {
            return w.type === deckType;
          })[0];
          common.addCount(winLoss, m);
        }
      });
      winLosss.forEach((w: any) => {
        w.games = w.win + w.lose;
        w.percentage = w.win / (w.win + w.lose);
        let o: any = {params: null};
        if (w.tempType === '') {
          o.params = Object.assign({}, options).params.append('race', w.race);
        } else {
          o.params = Object.assign({}, options).params.append('type', w.tempType);
        }
        this.http.get<any>(this.baseUrl + '/api/v1/tournament', o).subscribe((r: any) => {
          w.users = r.length;
          this.total += r.length;
        });
      });
      this.winLosss = winLosss;
      this.sortedData = this.winLosss.slice();
      this.dataSource = new MatTableDataSource(this.winLosss);
      this.dataSource.sort = this.sort;
    });
  }

  move(element) {
    const logs = [];
    this.matchUpLogs.forEach((m: MatchUpLog) => {
      if (element.type) {
        if (element.type === this.getPlayerType(m)) {
          logs.push(m);
        }
      } else {
        if (element.type === this.getDeckType(m)) {
          logs.push(m);
        }
      }
    });
    this.moveDetail.emit({
      type: element.type,
      process: [],
      logs: logs
    });
  }

  private getDeckType(m: MatchUpLog): string {
    let deckType = '';
    if (m.PlayerType !== '' && m.PlayerRace !== '') {
      const str = common.isMobile() ? '\n' : '';
      deckType = this.getPlayerType(m) + str + '(' + m.PlayerRace + ')';
    } else if (m.PlayerType !== '') {
      deckType = this.getPlayerType(m);
    } else if (m.PlayerRace !== '') {
      deckType = '(' + m.PlayerRace + ')';
    } else {
      deckType = 'その他';
    }
    return deckType;
  }

  private getPlayerType(m: MatchUpLog): string {
    return m.PlayerType.replace('..', '');
  }
}
