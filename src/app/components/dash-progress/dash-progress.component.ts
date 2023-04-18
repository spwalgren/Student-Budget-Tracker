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

const categorySpentTotals: { [key: string]: CategoryTotals } = {};
const categoryBudgetTotals: { [key: string]: CategoryTotals } = {};
const prevCategoryTotals: { [key: string]: CategoryTotals } = {};
const prevBudgCategoryTotals: { [key: string]: CategoryTotals } = {};


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
  weeklyTotalBudget: number = 0;
  monthlyTotalSpent: number = 0;
  monthlyTotalBudget: number = 0;
  yearlyTotalSpent: number = 0;
  yearlyTotalBudget: number = 0;
  prevWeeklyTotalSpent: number = 0;
  prevWeeklyTotalBudget: number = 0;
  prevMonthlyTotalSpent: number = 0;
  prevMonthlyTotalBudget: number = 0;
  prevYearlyTotalSpent: number = 0;
  prevYearlyTotalBudget: number = 0;

  subWeekly: number = 0;
  subMonthly: number = 0;
  subYearly: number = 0;
  subPrevWeekly: number = 0;
  subPrevMonthly: number = 0;
  subPrevYearly: number = 0;

  weeklyCategories: string[] = [];
  monthlyCategories: string[] = [];
  yearlyCategories: string[] = [];

  constructor( private progressService: ProgressService) { };

  ngOnInit() {
    this.progressService.GetProgress().subscribe((progress) => {
      if (!progress.err) {
        progress.data.forEach(elem => {
          const category = elem.category;

          if (!categorySpentTotals[category]) {
            categorySpentTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
            categoryBudgetTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
            if (elem.frequency == Period.weekly) {
              this.weeklyCategories.push(category);
            } else if (elem.frequency == Period.monthly){
              this.monthlyCategories.push(category);
            } else {
              this.yearlyCategories.push(category);
            }
          }
          if (elem.frequency == Period.weekly) {
            categorySpentTotals[category].weekly += elem.totalSpent;
            categoryBudgetTotals[category].weekly += elem.totalSpent;
            this.weeklyTotalSpent += elem.totalSpent;
            this.weeklyTotalBudget += elem.budgetGoal;
          }
          if (elem.frequency == Period.monthly) {
            categorySpentTotals[category].monthly += elem.totalSpent;
            categoryBudgetTotals[category].monthly += elem.totalSpent;
            this.monthlyTotalSpent += elem.totalSpent;
            this.monthlyTotalBudget += elem.budgetGoal;
          }
          if (elem.frequency == Period.yearly) {
            categorySpentTotals[category].yearly += elem.totalSpent;
            categoryBudgetTotals[category].yearly += elem.totalSpent;
            this.yearlyTotalSpent += elem.totalSpent;
            this.yearlyTotalBudget += elem.budgetGoal;
          }
        });
      }
      if(this.weeklyTotalBudget - this.weeklyTotalSpent < 0){
        this.subWeekly = 0;
      }else {
        this.subWeekly = this.weeklyTotalBudget - this.weeklyTotalSpent;
      } 
      if(this.monthlyTotalBudget - this.monthlyTotalSpent < 0){
        this.subMonthly = 0;
      }else {
        this.subMonthly = this.monthlyTotalBudget - this.monthlyTotalSpent;
      } 
      if(this.yearlyTotalBudget - this.yearlyTotalSpent < 0){
        this.subYearly = 0;
      }else {
        this.subYearly = this.yearlyTotalBudget - this.yearlyTotalSpent;
      } 
    });
  
    // this.progressService.GetPreviousProgress().subscribe((progress) => {
    //   if (!progress.err) {
    //     progress.data.forEach(elem => {
    //       const category = elem.category;
    //       if (!prevCategoryTotals[category]) {
    //         prevCategoryTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
    //         prevBudgCategoryTotals[category] = { weekly: 0, monthly: 0, yearly: 0 };
    //       }
    //       if (elem.frequency == Period.weekly) {
    //         prevCategoryTotals[category].weekly += elem.totalSpent;
    //         prevBudgCategoryTotals[category].weekly += elem.totalSpent;
    //         this.prevWeeklyTotalSpent += elem.totalSpent;
    //         this.prevWeeklyTotalBudget += elem.totalSpent;
    //       }
    //       if (elem.frequency == Period.monthly) {
    //         prevCategoryTotals[category].monthly += elem.totalSpent;
    //         prevBudgCategoryTotals[category].monthly += elem.totalSpent;
    //         this.prevMonthlyTotalSpent += elem.totalSpent;
    //         this.prevMonthlyTotalBudget += elem.totalSpent;
    //       }
    //       if (elem.frequency == Period.yearly) {
    //         prevCategoryTotals[category].yearly += elem.totalSpent;
    //         prevBudgCategoryTotals[category].yearly += elem.totalSpent;
    //         this.prevYearlyTotalSpent += elem.totalSpent;
    //         this.prevYearlyTotalBudget += elem.totalSpent;
    //       }
    //     });
    //   }

    //   if(this.prevWeeklyTotalBudget - this.prevWeeklyTotalSpent < 0){
    //     this.subPrevWeekly = 0;
    //   }else {
    //     this.subPrevWeekly = this.prevWeeklyTotalBudget - this.prevWeeklyTotalSpent;
    //   } 
    //   if(this.prevMonthlyTotalBudget - this.prevMonthlyTotalSpent < 0){
    //     this.subPrevMonthly = 0;
    //   }else {
    //     this.subPrevMonthly = this.monthlyTotalBudget - this.monthlyTotalSpent;
    //   } 
    //   if(this.prevYearlyTotalBudget - this.prevYearlyTotalSpent < 0){
    //     this.subPrevYearly = 0;
    //   }else {
    //     this.subPrevYearly = this.prevYearlyTotalBudget - this.prevYearlyTotalSpent;
    //   } 
  
    // });

    
  

  }
  
}
