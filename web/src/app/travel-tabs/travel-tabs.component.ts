import { Component, OnInit } from '@angular/core';
import { ApiService, Travel } from "../api.service";
import { map, shareReplay } from "rxjs/operators";
import { Observable } from "rxjs";
import { MatDialog } from "@angular/material/dialog";
import { NewComponent } from "../new/new.component";

interface TabData {
  title: string;
  icon: string;
  data?: Observable<Travel[]>
}

@Component({
  selector: 'app-travel-tabs',
  templateUrl: './travel-tabs.component.html',
  styleUrls: ['./travel-tabs.component.scss']
})

export class TravelTabsComponent implements OnInit {
  tabs: TabData[] = [
    {
      title: 'Upcoming',
      icon: 'flight_takeoff',
    },
    {
      title: 'Done',
      icon: 'flight_land',
    }
  ];

  constructor(
    private apiService: ApiService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit(): void {
    this.refresh();
  }

  refresh() {
    const api$ = this.apiService.getTravels$().pipe(shareReplay(1));

    this.tabs = [...this.tabs.map(x => ({
      ...x,
      data: api$.pipe(map(travel => travel[x.title.toLowerCase()]))
    }))];
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(NewComponent, {
      width: '250px',
      data: {},
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        void this.apiService.createTravel$(result).toPromise();
        this.refresh();
      }
    });
  }

}
