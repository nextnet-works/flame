from backend.local import LocalBackend
from backend.mqtt import MqttBackend
from backend.p2p import PointToPointBackend
from object_factory import ObjectFactory


class BackendProvider(ObjectFactory):
    def get(self, backend_name, **kargs):
        return self.create(backend_name, **kargs)


provider = BackendProvider()
provider.register('local', LocalBackend())
provider.register('p2p', PointToPointBackend())
provider.register('mqtt', MqttBackend())
