import { Component, Input } from '@angular/core';
import { EventContent } from 'src/types/calendar-system';

@Component({
  selector: 'app-event-card',
  templateUrl: './event-card.component.html',
  styleUrls: ['./event-card.component.css']
})
export class EventCardComponent {

  @Input()
  eventData?: EventContent;

}
