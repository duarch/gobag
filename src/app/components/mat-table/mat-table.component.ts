import { NgFor } from '@angular/common';
import { Component } from '@angular/core';



@Component({
  selector: 'app-mat-table',
  standalone: true,
  imports: [NgFor],
  templateUrl: './mat-table.component.html',
  styleUrl: './mat-table.component.css'
})



export class MatTableComponent {
  dados = [
    { id: 1, nome: 'Kit de primeiros socorros', descricao: 'Contém itens essenciais de primeiros socorros', peso: '1kg', medidas: '20cm x 10cm x 10cm', dataValidade: '2025-12-31' },
    { id: 2, nome: 'Água potável', descricao: 'Garrafas de água potável', peso: '500g', medidas: '30cm x 10cm x 10cm', dataValidade: '2025-12-31' },
    { id: 3, nome: 'Comida enlatada', descricao: 'Alimentos não perecíveis', peso: '1kg', medidas: '15cm x 10cm x 10cm', dataValidade: '2025-12-31' },
    { id: 4, nome: 'Lanterna', descricao: 'Lanterna alimentada por manivela', peso: '250g', medidas: '15cm x 5cm x 5cm', dataValidade: 'N/A' },
    { id: 5, nome: 'Cobertor de emergência', descricao: 'Cobertor térmico para manter o calor corporal', peso: '150g', medidas: '10cm x 5cm x 5cm', dataValidade: 'N/A' },
    // Adicione mais dados aqui conforme necessário
  ];
  
}



