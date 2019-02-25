import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  entryComponents: []
})
export class AppComponent implements OnInit {

  type;
  logs;
  deckType;

  ngOnInit(): void {
    this.type = 'search';
  }

  setType(type: string, data) {
    window.scrollTo(0, 0);
    this.type = type;
    this.deckType = data.type;
    this.logs = data.logs;
  }

}
