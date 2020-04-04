import { Component, OnInit } from '@angular/core';
import { ApiService, Travel } from "../api.service";
import { map, shareReplay } from "rxjs/operators";
import { Observable } from "rxjs";
import { MatDialog } from "@angular/material/dialog";
import { NewComponent } from "../new/new.component";

interface TabData {
  title: string;
  icon: string;
  data: Observable<Travel[]>
}

@Component({
  selector: 'app-travel-tabs',
  templateUrl: './travel-tabs.component.html',
  styleUrls: ['./travel-tabs.component.scss']
})

export class TravelTabsComponent implements OnInit {
  tabs: TabData[] = [];

  constructor(
    private apiService: ApiService,
    public dialog: MatDialog,
  ) {
  }

  ngOnInit(): void {
    this.setup();
  }

  refresh() {
    this.setup();
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

  setup() {
    const api$ = this.apiService.getTravels$().pipe(shareReplay(1));

    const upcoming$ = api$.pipe(map(x => x.upcoming));
    const done$ = api$.pipe(map(x => x.done));

    this.tabs = [
      {
        title: 'Upcoming',
        icon: 'flight_takeoff',
        data: upcoming$,
      },
      {
        title: 'Done',
        icon: 'flight_land',
        data: done$,
      }
    ];

  }

}
