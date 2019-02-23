import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { MatchUpLog } from '../../models/match-up-log';
import { common } from '../../common';
import { MatTableDataSource, MatSort } from '@angular/material';

@Component({
  selector: 'app-components-detail',
  templateUrl: './detail.component.html',
  styleUrls: ['./detail.component.css'],
})

export class AppComponentsDetailComponent implements OnInit {

  @Input() deckType: string;
  @Input() matchUpLogs: MatchUpLog[] = [];
  @ViewChild(MatSort) sort: MatSort;

  displayedColumns = ['type', 'win', 'lose', 'percentage'];
  dataSource;
  winLosss: any[];
  sortedData: any[];
  useCards: string[] = [];
  urls: string[] = [];

  ngOnInit() {
    this.matchUpLogs.forEach((log: MatchUpLog) => {
      this.setCards(log.PlayerUseCards);
      this.setUrls(log.Url);
    });
    this.hoge(this.matchUpLogs);
  }

  hoge(logs: MatchUpLog[]) {
    const types: string[] = [];
    const winLosss: {[key: string]: any}[] = [];

    logs.forEach((m: MatchUpLog) => {
      const deckType = this.getOpponentDeckType(m);
      if (!types.includes(deckType)) {
        types.push(deckType);
        const w = {type: deckType, win: 0, lose: 0};
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
    });
    this.winLosss = winLosss;
    this.sortedData = this.winLosss.slice();
    this.dataSource = new MatTableDataSource(this.winLosss);
    this.dataSource.sort = this.sort;
  }

  setUrls(url: string) {
    if (!this.urls.includes(url)) {
      this.urls.push(url);
    }
  }

  setCards(cards: string[]) {
    cards.forEach((c: string) => {
      if (!this.useCards.includes(c)) {
        this.useCards.push(c);
      }
    });
    this.useCards.sort();
  }

  private getOpponentDeckType(m: MatchUpLog): string {
    let deckType = '';
    if (m.OpponentType !== '' && m.OpponentRace !== '') {
      const str = common.isMobile() ? '\n' : '';
      deckType = m.OpponentType + str + '(' + m.OpponentRace + ')';
    } else if (m.OpponentType !== '') {
      deckType = m.OpponentType;
    } else if (m.OpponentRace !== '') {
      deckType = '(' + m.OpponentRace + ')';
    } else {
      deckType = 'その他';
    }
    return deckType;
  }
}
