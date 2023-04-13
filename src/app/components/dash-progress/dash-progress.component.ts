import { Component } from '@angular/core';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatTabsModule } from '@angular/material/tabs';
import { Progress } from 'src/types/progress.system';
import { ProgressService } from 'src/app/progress.service';
import { Period } from 'src/types/budget-system';
import { Observable, catchError, of, map, from } from 'rxjs';
import { Injectable } from '@angular/core';



@Component({
  selector: 'app-dash-progress',
  templateUrl: './dash-progress.component.html',
  styleUrls: ['./dash-progress.component.css']
})
export class DashProgressComponent {
  tab: any;
  progress: Progress[] = [];
  freq: [Period.weekly, Period.monthly, Period.yearly] | undefined;
  weeklyTotalSpent: number = 0;
  monthlyTotalSpent: number = 0;
  yearlyTotalSpent: number = 0;

  constructor( private progressService: ProgressService) { };

  ngOnInit() {
      this.progressService.GetProgress().subscribe((progress) => {
        //this.progress = data;
        if (!progress.err){
          progress.data.forEach(elem => {
            if ( elem.frequency == Period.weekly){
              this.weeklyTotalSpent += elem.totalSpent;
            }
            if ( elem.frequency == Period.monthly){
              this.monthlyTotalSpent += elem.totalSpent;
            }
            if ( elem.frequency == Period.yearly){
              this.yearlyTotalSpent += elem.totalSpent;
            }
        
          });
          }
      });
    }
  
}
