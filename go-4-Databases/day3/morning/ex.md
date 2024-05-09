1. uma colecao
2. 25359 documentos
3. 278.5 KB

4. {
  returnType: 'this',
  serverVersions: [ '0.0.0', '999.999.999' ],
  apiVersions: [ 0, Infinity ],
  topologies: [ 'ReplSet', 'Sharded', 'LoadBalanced', 'Standalone' ],
  returnsPromise: false,
  deprecated: false,
  platforms: [ 'Compass', 'Browser', 'CLI' ],
  isDirectShellCommand: false,
  acceptsRawInput: false,
  shellCommandCompleter: undefined,
  help: [Function (anonymous)] Help
}

EX 1
client.db('sample_restaurante').collection('restaurant').find({restaurante_id: "40364681"}, {  
    _id: 0, 
    barrio: 1, 
    restaurante_id: 1, 
    nombre: 1, 
    tipo_cocina: 1 });

EX 2
client.db('sample_restaurante').collection('restaurant').find({nombre: {$regex: /bake/gm}}, {  
    _id: 0, 
    barrio: 1, 
    restaurante_id: 1, 
    nombre: 1, 
    tipo_cocina: 1 });

EX 3

db.restaurant.aggregate([
  { $match: { barrio: "Bronx" } },
  { $group: { _id: "$tipo_cocina", count: { $sum: 1 } } }
])

EX 4

db.restaurantes.find({
  grados: {
    
  }})

  EX 5

  db.restaurant.find({
  $or: [
    { "direction.coord": { $eq: {} } },
    { "direction.coord": { $eq: 0 } }
  ]})

  EX 6

  db.restaurant.find({
  $or: [
    { "direction.coord": { $eq: {} } },
    { "direction.coord": { $eq: 0 } }
  ]}, {  
    _id: 0, 
    barrio: 1, 
    restaurante_id: 1, 
    nombre: 1, 
    tipo_cocina: 1 })

EX 7 
    db.restaurant.aggregate([
        { $group: { _id: "$tipo_cocina", count: { $sum: 1 } } },

        {$sort: {
            count: -1
        }},
            { $limit: 3}
        ])

ex 8

db.restaurant.aggregate([
  { $addFields: { 
      tamanhoGrados: { $size: "$grados" } 
  }},
  { $match: { 
      tamanhoGrados: { $gte: 3 } 
  }},
  {$unwind: "$grados"},
  { $group: { 
      _id: "$nombre",
      avarege: { $avg: "$grados.puntaje" } 
  }},
  {$sort: {
    avarege: -1}
  }
])

EX 9

db.restaurant.createIndex({ "direccion.coord": "2dsphere" })
db.restaurant.find({
  "direccion.coord": {
    $near: {
      $geometry: {
        type: "Point",
        coordinates: [-73.93414657, 40.82302903]
      },
      $maxDistance: 500
    }
  }
})