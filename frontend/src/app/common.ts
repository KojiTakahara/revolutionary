import { MatchUpLog } from './models';

export class Common {
  isMobile(): boolean {
    const ua = navigator.userAgent.toLowerCase();
    const sdev = /iphone;|(android|nokia|blackberry|bb10;).+mobile|android.+fennec|opera.+mobi|windows phone|symbianos/;
    return sdev.test(ua);
  }

  addCount(winLoss: {[key: string]: any}, m: MatchUpLog) {
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

export const common: Common = new Common();
