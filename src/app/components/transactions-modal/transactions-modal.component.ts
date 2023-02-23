import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog'

@Component({
  selector: 'app-transactions-modal',
  templateUrl: './transactions-modal.component.html',
  styleUrls: ['./transactions-modal.component.css']
})
export class TransactionsModalComponent {
  constructor(public dialogRef: MatDialogRef<TransactionsModalComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any) { 

    } //data b/w data and source

    ngOnInit(){

    }

    save(){
      this.dialogRef.close("DATA SAVED");
      //save doc
    }
}
