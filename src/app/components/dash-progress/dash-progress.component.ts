import { Component } from '@angular/core';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatTabsModule } from '@angular/material/tabs';
import { Progress } from 'src/types/progress.system';
import { ProgressService } from 'src/app/progress.service';
import { Period } from 'src/types/budget-system';
import { Observable, catchError, of, map, from } from 'rxjs';
import { Injectable } from '@angular/core';

export interface CategoryTotals {
  weekly: number;
  monthly: number;
  yearly: number;
}

const categoryTotals: { [key: string]: CategoryTotals } = {};
const prevCategoryTotals: { [key: string]: CategoryTotals } = {};

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
  prevWeeklyTotalSpent: number = 0;
  prevMonthlyTotalSpent: number = 0;
  prevYearlyTotalSpent: number = 0;
 

  constructor( private progressService: ProgressService) { };

  ngOnInit() {
    this.progressService.GetProgress().subscribe((progress) => {
      if (!progress.err) {
        progress.data.forEach(elem => {
          const category = elem.category;
          if (!categoryTotals[category]) {
            categoryTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
          }
          if (elem.frequency == Period.weekly) {
            categoryTotals[category].weekly += elem.totalSpent;
            this.weeklyTotalSpent += elem.totalSpent;
          }
          if (elem.frequency == Period.monthly) {
            categoryTotals[category].monthly += elem.totalSpent;
            this.monthlyTotalSpent += elem.totalSpent;
          }
          if (elem.frequency == Period.yearly) {
            categoryTotals[category].yearly += elem.totalSpent;
            this.yearlyTotalSpent += elem.totalSpent;
          }
        });
      }
    });

    this.progressService.GetPreviousProgress().subscribe((progress) => {
      if (!progress.err) {
        progress.data.forEach(elem => {
          const category = elem.category;
          if (!prevCategoryTotals[category]) {
            prevCategoryTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
          }
          if (elem.frequency == Period.weekly) {
            prevCategoryTotals[category].weekly += elem.totalSpent;
            this.prevWeeklyTotalSpent += elem.totalSpent;
          }
          if (elem.frequency == Period.monthly) {
            prevCategoryTotals[category].monthly += elem.totalSpent;
            this.prevMonthlyTotalSpent += elem.totalSpent;
          }
          if (elem.frequency == Period.yearly) {
            prevCategoryTotals[category].yearly += elem.totalSpent;
            this.prevYearlyTotalSpent += elem.totalSpent;
          }
        });
      }
    });


  }
  
}
