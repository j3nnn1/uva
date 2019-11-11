import {Component, OnInit} from '@angular/core';
import * as SampleJson from '../assets/timeseries/uva_usd.json';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent implements OnInit {
  // this contain json data
  public multi: any;
  // this contain json data

  title = 'Uva';
  jsonUrl = 'assets/timeseries/uva_usd2.json';
  view: any[] = [1300, 400];

  // options for the chart
  showXAxis = true;
  showYAxis = true;
  gradient = false;
  showLegend = true;
  showXAxisLabel = true;
  xAxisLabel = 'Tiempo';
  showYAxisLabel = true;
  yAxisLabel = 'Pesos';
  timeline = true;
  colorScheme = {
    domain: ['#9370DB', '#87CEFA', '#FA8072', '#FF7F50', '#90EE90', '#9370DB']
  };
  // pie
  showLabels = true;
  // options for the chart

  ngOnInit() {
    this.multi = SampleJson.default;
    console.log('onInit', this.multi.default);
  }
  onSelect(something: any) {
      console.log('on select graph');
  }
}
