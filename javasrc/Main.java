package javasrc;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import spoon.Launcher;
import spoon.reflect.*;
import spoon.reflect.declaration.*;
import spoon.reflect.reference.*;
import spoon.support.reflect.declaration.*;
import spoon.support.reflect.reference.*;
import spoon.reflect.factory.ModuleFactory;
import java.io.File;
import java.util.concurrent.locks.ReentrantLock;

class Main {
    public static void main(String[] args) {

        String source = args[0];

        Launcher launcher = new Launcher();
        launcher.addInputResource(source);
        CtModel model = launcher.buildModel();

        var root = model.getAllModules().iterator().next().getRootPackage();
        if ( root == null ) {
            throw new RuntimeException("null module");
        }

        Gson gson = new GsonBuilder()
            .setPrettyPrinting()
            /*
            .registerTypeAdapter(CtReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtActualTypeContainer.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtArrayTypeReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtCatchVariableReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtExecutableReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtFieldReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtIntersectionTypeReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtLocalVariableReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtModuleReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtPackageReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtParameterReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeMemberWildcardImportReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeParameterReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtUnboundVariableReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtVariableReference.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtWildcardReference.class, new CtReferenceJSONSerializer())

            .registerTypeAdapter(CtArrayTypeReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtCatchVariableReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtExecutableReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtFieldReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtIntersectionTypeReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtLocalVariableReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtModuleReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtPackageReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtParameterReferenceImpl .class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeMemberWildcardImportReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeParameterReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtTypeReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtUnboundVariableReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtVariableReferenceImpl.class, new CtReferenceJSONSerializer())
            .registerTypeAdapter(CtWildcardReferenceImpl.class, new CtReferenceJSONSerializer())
            */



            .registerTypeAdapter(ModuleFactory.CtUnnamedModule.class, new NullJSONSerializer())
            .registerTypeAdapter(File.class, new NullJSONSerializer())
            .registerTypeAdapter(ReentrantLock.class, new NullJSONSerializer())

            /*
            .registerTypeAdapter(CtPackage.class, new CtPackageJSONSerializer())
            .registerTypeAdapter(CtPackageImpl.class, new CtPackageJSONSerializer())
            .registerTypeAdapter(CtType.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtAnnotationTypeImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtClassImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtEnumImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtInterfaceImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtRecordImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtTypeImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtTypeParameterImpl.class, new CtTypeJsonSerializer())
            .registerTypeAdapter(CtField.class, new CtFieldJsonSerializer())
            .registerTypeAdapter(CtEnumValueImpl.class, new CtFieldJsonSerializer())
            .registerTypeAdapter(CtFieldImpl.class, new CtFieldJsonSerializer())
            */

            .registerTypeAdapter(CtElementImpl.class, new CtElementJsonSerializer())

            .create();
        
        String js = gson.toJson(root);
        System.out.println(js);
    }

}