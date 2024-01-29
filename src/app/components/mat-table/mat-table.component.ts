import { Component } from '@angular/core';

@Component({
  selector: 'app-mat-table',
  standalone: true,
  imports: [],
  template: `
  <table mat-table [dataSource]="items" class="mat-elevation-z8">
    <!-- ID Column -->
    <ng-container matColumnDef="id">
      <th mat-header-cell *matHeaderCellDef> ID </th>
      <td mat-cell *matCellDef="let item"> {{item.id}} </td>
    </ng-container>

    <!-- Nome Column -->
    <ng-container matColumnDef="nome">
      <th mat-header-cell *matHeaderCellDef> Nome </th>
      <td mat-cell *matCellDef="let item"> {{item.nome}} </td>
    </ng-container>

    <!-- Descrição Column -->
    <ng-container matColumnDef="descricao">
      <th mat-header-cell *matHeaderCellDef> Descrição </th>
      <td mat-cell *matCellDef="let item"> {{item.descricao}} </td>
    </ng-container>

    <!-- Peso Column -->
    <ng-container matColumnDef="peso">
      <th mat-header-cell *matHeaderCellDef> Peso </th>
      <td mat-cell *matCellDef="let item"> {{item.peso}} </td>
    </ng-container>

    <!-- Medidas Column -->
    <ng-container matColumnDef="medidas">
      <th mat-header-cell *matHeaderCellDef> Medidas </th>
      <td mat-cell *matCellDef="let item"> {{item.medidas}} </td>
    </ng-container>

    <!-- Data de validade Column -->
    <ng-container matColumnDef="dataValidade">
      <th mat-header-cell *matHeaderCellDef> Data de validade </th>
      <td mat-cell *matCellDef="let item"> {{item.dataValidade}} </td>
    </ng-container>

    <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
    <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
  </table>
`,
  styleUrl: './mat-table.component.css'
})
export class MatTableComponent {
  items = ELEMENT_DATA;
  displayedColumns: string[] = ['id', 'nome', 'descricao', 'peso', 'medidas', 'dataValidade'];
}


const ELEMENT_DATA = [
  // Adicione seus dados aqui. Por exemplo:
  {id: 1, nome: 'Item 1', descricao: 'Descrição 1', peso: 1.0079, medidas: 'Medida 1', dataValidade: '20/12/2023'},
];


