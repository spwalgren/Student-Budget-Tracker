import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css']
})
export class AlertComponent {

  @Input()
  type?: "success" | "info" | "error" | undefined
  @Input()
  message?: string;

  alertClassName(): string {
    if (this.type)
      return `alert alert--${this.type}`;
    else
      return 'alert';
  }

  fontIcon(): string {
    if (this.type == "success")
      return "task_alt";
    else if (this.type == "info")
      return "info_outline";
    else if (this.type == "error")
      return "error_outline";
    else
      return "error_outline";
  }
}
