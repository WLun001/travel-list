import { Component, EventEmitter, Input, OnDestroy, Output } from '@angular/core';
import { ApiService, Travel } from "../api.service";
import { Subscription } from "rxjs";

@Component({
  selector: 'app-travel-list',
  templateUrl: './travel-list.component.html',
  styleUrls: ['./travel-list.component.scss']
})
export class TravelListComponent implements OnDestroy {
  @Input() travels: Travel[];
  @Output() refresh = new EventEmitter();
  private subscription = new Subscription();

  constructor(private apiService: ApiService) {
  }

  markAsDone(id: string) {
    const travel = this.travels.find(x => x.id === id);
    travel.done = true;
    this.subscription = this.apiService.updateTravel$(travel)
      .subscribe(() => this.refresh.emit());
  }

  delete(id: string) {
    this.subscription = this.apiService.deleteTravel$(id)
      .subscribe(() => this.refresh.emit());
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

}
