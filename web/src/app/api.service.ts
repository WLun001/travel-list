import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { environment } from "../environments/environment";
import { filter, map } from "rxjs/operators";


export interface Travel {
  id: string;
  name: string;
  photo: string;
  done: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) {
  }

  getTravels$() {
    return this.http.get<Travel[]>(`${ environment.apiUrl }/travels`)
      .pipe(
        filter(x => x?.length > 0),
        map(x => ({
          upcoming: x.filter(y => !y.done),
          done: x.filter(y => y.done)
        })),
      );
  }

  createTravel$(travel: Travel) {
    return this.http.post(`${ environment.apiUrl }/travels`, travel);
  }

  updateTravel$(travel: Travel) {
    return this.http.put(`${ environment.apiUrl }/travels/${ travel.id }`, travel);
  }

  deleteTravel$(id: string) {
    return this.http.delete(`${ environment.apiUrl }/travels/${ id }`);
  }
}
